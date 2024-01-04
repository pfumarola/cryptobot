package main

import (
	"context"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	binance_connector "github.com/binance/binance-connector-go"
	"github.com/joho/godotenv"
	"github.com/pfumarola/cryptopbot/internal"
)

type Trader struct {
	client            *binance_connector.Client
	firstPair         string
	secondPair        string
	thirdPair         string
	holdingAsset      string
	minProfitability  float64
	feePercentage     float64
	orderAmount       float64
	bookTickerService *binance_connector.TickerBookTicker
	bookTick          []*binance_connector.TickerBookTickerResponse
	exchangeInfo      *internal.ExchangeInfo
	telegramCh        chan string
}

func NewTrader(client *binance_connector.Client, holdingAsset string, minProfitability float64, feePercentage float64, orderAmount float64, telegramCh chan string) *Trader {
	exchangeInfo := new(internal.ExchangeInfo)
	return &Trader{
		holdingAsset:      holdingAsset,
		minProfitability:  minProfitability,
		feePercentage:     feePercentage,
		orderAmount:       orderAmount,
		client:            client,
		bookTickerService: client.NewTickerBookTickerService(),
		exchangeInfo:      exchangeInfo.NewExchangeInfo(client),
		telegramCh:        telegramCh,
	}
}

func (t *Trader) checkOpportunity() (float64, float64) {
	bid1, ask1 := t.getPrice(t.firstPair)
	bid2, ask2 := t.getPrice(t.secondPair)
	bid3, ask3 := t.getPrice(t.thirdPair)
	potentialProfitBSS := (1/ask1)*bid2*bid3 - 1 - (t.feePercentage * 4)
	potentialProfitBBS := (1/ask3)*(1/ask2)*bid1 - 1 - (t.feePercentage * 4)
	if math.IsInf(potentialProfitBSS, 0) {
		potentialProfitBSS = 0.0
	}
	if math.IsInf(potentialProfitBBS, 0) {
		potentialProfitBBS = 0.0
	}
	//fmt.Println(fmt.Sprintf("BBS %.4f %.4f %.14f %.4f", potentialProfitBBS, ask1, bid2, bid3))
	if potentialProfitBSS > t.minProfitability {
		s := fmt.Sprintf("BSS %s->%s->%s %.4f %.4f %.14f %.4f", t.firstPair, t.secondPair, t.thirdPair, potentialProfitBSS, ask1, bid2, bid3)
		fmt.Println(s)
		t.telegramCh <- s
		t.createOrder(t.firstPair, t.secondPair, t.thirdPair, "BUY", t.orderAmount)
	} else if potentialProfitBBS > t.minProfitability {
		//fmt.Println("Opportunity found (BUY BUY SELL))", potentialProfitBBS)
		//t.createOrder(t.firstPair, t.thirdPair, t.secondPair, "BUY", 100, "BBS")
	}
	return potentialProfitBSS, potentialProfitBBS
}

func (t *Trader) updateBookTick() {
	var err error
	t.bookTick, err = t.bookTickerService.Do(context.Background())
	t.check(err)
}

func (t *Trader) onTick(ticker *time.Ticker, done chan bool) {
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			go t.updateBookTick()
			//cycle through all the pairs in ArbitrageTriplets
			maxPotentialProfitBSS := -100000.0
			maxPotentialProfitBBS := -100000.0
			for _, v := range internal.ArbitrageTriplets {
				t.firstPair = v.FirstPair
				t.secondPair = v.SecondPair
				t.thirdPair = v.ThirdPair
				potentialProfitBSS, potentialProfitBBS := t.checkOpportunity()
				if potentialProfitBSS > maxPotentialProfitBSS {
					maxPotentialProfitBSS = potentialProfitBSS
				}
				if potentialProfitBBS > maxPotentialProfitBBS {
					maxPotentialProfitBBS = potentialProfitBBS
				}
			}
			//timestamp := time.Now().Format("2006-01-02 15:04:05")
			//fmt.Println(timestamp, " Max potential profit: ", fmt.Sprintf("%.8f", maxPotentialProfitBSS), " ", fmt.Sprintf("%.8f", maxPotentialProfitBBS))
		}
	}
}

func (t *Trader) ticker() {
	ticker := time.NewTicker(time.Second * 1)
	done := make(chan bool)
	go t.onTick(ticker, done)
}

func (t *Trader) getPrice(symbol string) (float64, float64) {
	symbol = strings.ReplaceAll(symbol, "-", "")
	for _, v := range t.bookTick {
		if v.Symbol == symbol {
			bid, _ := strconv.ParseFloat(v.BidPrice, 32)
			ask, _ := strconv.ParseFloat(v.AskPrice, 32)
			return bid, ask
		}
	}
	return 0.0, 0.0
}

func (t *Trader) getBalance(s1 string, s2 string, s3 string) []binance_connector.Balance {
	val, err := t.client.NewGetAccountService().Do(context.Background())
	t.check(err)
	return t.filterBalance(val.Balances, []string{s1, s2, s3, "BNB"})
}

func (t *Trader) filterBalance(balances []binance_connector.Balance, names []string) []binance_connector.Balance {
	var filtered []binance_connector.Balance
	for _, balance := range balances {
		for _, name := range names {
			if balance.Asset == name {
				filtered = append(filtered, balance)
				break
			}
		}
	}
	return filtered
}

func (t *Trader) splitPair(pair string) (string, string) {
	split := strings.Split(pair, "-")
	return split[0], split[1]
}

func (t *Trader) createOrder(symbol1, symbol2, symbol3, side1 string, quantity1 float64) {

	base1, quote1 := t.splitPair(symbol1)
	base2, quote2 := t.splitPair(symbol2)
	base3, quote3 := t.splitPair(symbol3)

	tgMsg := ""

	scale := 0.0
	var balance []binance_connector.Balance

	base1Decimals := countDecimals(t.exchangeInfo.GetSymbolLotSize(base1 + quote1).StepSize)
	base2Decimals := countDecimals(t.exchangeInfo.GetSymbolLotSize(base2 + quote2).StepSize)
	base3Decimals := countDecimals(t.exchangeInfo.GetSymbolLotSize(base3 + quote3).StepSize)
	fmt.Println("base3Decimals: ", base3Decimals)
	_ = base1Decimals

	balance = t.getBalance(base1, quote1, quote2)
	fmt.Println("Balance: ", balance)
	tgMsg += fmt.Sprintf("Balance: %s \n", balance)

	newOrder, err := t.client.NewCreateOrderService().Symbol(base1 + quote1).
		Side("BUY").Type("MARKET").QuoteOrderQty(quantity1).
		Do(context.Background())
	t.check(err)
	fmt.Println(binance_connector.PrettyPrint(newOrder))
	tgMsg += fmt.Sprintf("Order: %s \n", binance_connector.PrettyPrint(newOrder))

	t.getBalance(base1, quote1, quote2)
	fmt.Println("Balance: ", balance)
	tgMsg += fmt.Sprintf("Balance: %s \n", balance)

	quantity2, _ := strconv.ParseFloat(newOrder.(*binance_connector.CreateOrderResponseFULL).ExecutedQty, 64)
	scale = math.Pow(10, float64(base2Decimals))
	floatQuantity2 := math.Floor(quantity2*scale) / scale

	t.check(err)
	fmt.Println("Quantity2: ", floatQuantity2)

	newOrder, err = t.client.NewCreateOrderService().Symbol(base2 + quote2).
		Side("SELL").Type("MARKET").Quantity(floatQuantity2).
		Do(context.Background())
	t.check(err)
	fmt.Println(binance_connector.PrettyPrint(newOrder))
	tgMsg += fmt.Sprintf("Order: %s \n", binance_connector.PrettyPrint(newOrder))

	t.getBalance(base1, quote1, quote2)
	fmt.Println("Balance: ", balance)
	tgMsg += fmt.Sprintf("Balance: %s \n", balance)

	quantity3, _ := strconv.ParseFloat(newOrder.(*binance_connector.CreateOrderResponseFULL).CumulativeQuoteQty, 64)
	scale = math.Pow(10, float64(base3Decimals))
	floatQuantity3 := math.Floor(quantity3*scale) / scale

	t.check(err)

	fmt.Println("Quantity3: ", floatQuantity3)
	newOrder, err = t.client.NewCreateOrderService().Symbol(base3 + quote3).
		Side("SELL").Type("MARKET").Quantity(floatQuantity3).
		Do(context.Background())
	t.check(err)
	fmt.Println(binance_connector.PrettyPrint(newOrder))
	tgMsg += fmt.Sprintf("Order: %s \n", binance_connector.PrettyPrint(newOrder))

	balance = t.getBalance(base1, quote1, quote2)
	fmt.Println("Balance: ", balance)
	tgMsg += fmt.Sprintf("Balance: %s \n", balance)

	fmt.Println("End of arbitrage")
	tgMsg += fmt.Sprintf("End of arbitrage")

	t.telegramCh <- tgMsg
}

func (t *Trader) check(e error) {
	if e != nil {
		t.telegramCh <- fmt.Sprintf("Error: %s", e)
		panic(e)
	}
}

func countDecimals(num string) int {
	// Find the index of the decimal point
	index := strings.Index(num, "1")
	if index == 0 {
		return 0
	}
	return index - 1
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// BINANCE_API_KEY := os.Getenv("TEST_BINANCE_API_KEY")
	// BINANCE_API_SECRET := os.Getenv("TEST_BINANCE_API_SECRET")
	// client := binance_connector.NewClient(BINANCE_API_KEY, BINANCE_API_SECRET, "https://testnet.binance.vision")

	BINANCE_API_KEY := os.Getenv("BINANCE_API_KEY")
	BINANCE_API_SECRET := os.Getenv("BINANCE_API_SECRET")
	client := binance_connector.NewClient(BINANCE_API_KEY, BINANCE_API_SECRET) //, "https://testnet.binance.vision") //"https://testnet.binance.vision"

	//client.Debug = true

	TELEGRAM_TOKEN := os.Getenv("TELEGRAM_TOKEN")
	TELEGRAM_CHAT_ID := os.Getenv("TELEGRAM_CHAT_ID")

	telegramCh := make(chan string)

	trader := NewTrader(client, "USDT", 0.01, 0.00075, 100, telegramCh) // 0.01 -> 1% min profitability

	go broker(TELEGRAM_TOKEN, TELEGRAM_CHAT_ID, telegramCh)

	trader.ticker()

	wg.Wait()
}

func broker(TELEGRAM_TOKEN string, TELEGRAM_CHAT_ID string, telegramCh chan string) {
	var telegramBot *internal.TelegramBot

	telegramBot = internal.NewTelegramBot(TELEGRAM_TOKEN)
	go telegramBot.Start()
	for {
		select {
		case msg := <-telegramCh:
			_ = msg
			telegramBot.SendMessage(TELEGRAM_CHAT_ID, msg)
		}
	}
}
