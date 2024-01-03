package mainold

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	binance_connector "github.com/binance/binance-connector-go"
	"github.com/joho/godotenv"
)

var firstPair = "TRXUSDT"
var secondPair = "TRXBTC"
var thirdPair = "BTCUSDT"

//TRX-USDT;TRX-BTC;BTC-USDT

var holdingAsset = "USDT"

// var holdingAsset = "USDT"
// var firstPair = "TRXUSDT"
// var secondPair = "TRXBTC"
// var thirdPair = "BTCUSDT"

var client *binance_connector.Client
var pairsCSV *os.File
var bookTick []*binance_connector.TickerBookTickerResponse

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkOpportunity() (float64, float64, float64, float64) {
	var err error
	bookTick, err = client.NewTickerBookTickerService().Do(context.Background())
	check(err)
	//fmt.Println(binance_connector.PrettyPrint(bookTick))
	_, ask1 := getPrice(firstPair)
	bid2, ask2 := getPrice(secondPair)
	bid3, _ := getPrice(thirdPair)
	//potentialProfitBBS := (1/ask1)*(1/ask2)*bid3 - 1
	potentialProfitBSS := (1/ask1)*bid2*bid3 - 1
	fmt.Println(fmt.Sprintf("%.4f %.4f %.14f %.4f", potentialProfitBSS, ask1, bid2, bid3))
	return potentialProfitBSS, ask1, ask2, bid3
}

func onTick(ticker *time.Ticker, done chan bool) {
	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			_ = t
			checkOpportunity()
			//getBalance()
			//createOrder(firstPair, secondPair, thirdPair, "BUY", 100)
		}
	}
}

func ticker() {
	//create a ticker that ticks every second
	ticker := time.NewTicker(time.Second)
	//create a channel to receive the ticks
	done := make(chan bool)
	//create a go routine that listens to the channel
	go onTick(ticker, done)
}

func getPrice(symbol string) (float64, float64) {
	// search and return latest price of a symbol in bookTick
	for _, v := range bookTick {
		if v.Symbol == symbol {
			//return v.BidPrice, v.AskPrice converted from string to float64
			bid, _ := strconv.ParseFloat(v.BidPrice, 32)
			ask, _ := strconv.ParseFloat(v.AskPrice, 32)
			return bid, ask
		}
	}
	return 0.0, 0.0
}

func getBalance() {
	val, err := client.NewGetAccountService().Do(context.Background())
	check(err)
	fmt.Println(filterBalance(val.Balances, []string{"USDT", "TRX", "ETH"}), "\n")
}

func filterBalance(balances []binance_connector.Balance, names []string) []binance_connector.Balance {
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

func createOrder(
	symbol1 string,
	symbol2 string,
	symbol3 string,
	side1 string,
	quantity1 float64,
) {
	// Create new order
	newOrder, err := client.NewCreateOrderService().Symbol(symbol1).
		Side("BUY").Type("MARKET").QuoteOrderQty(quantity1).
		Do(context.Background())
	check(err)
	fmt.Println(binance_connector.PrettyPrint(newOrder))

	getBalance()

	quantity2 := newOrder.(*binance_connector.CreateOrderResponseFULL).ExecutedQty
	floatQuantity2, _ := strconv.ParseFloat(quantity2, 64)
	check(err)

	// Create new order
	newOrder, err = client.NewCreateOrderService().Symbol(symbol2).
		Side("BUY").Type("MARKET").QuoteOrderQty(floatQuantity2).
		Do(context.Background())
	check(err)
	fmt.Println(binance_connector.PrettyPrint(newOrder))

	getBalance()

	quantity3 := newOrder.(*binance_connector.CreateOrderResponseFULL).ExecutedQty
	floatQuantity3, _ := strconv.ParseFloat(quantity3, 64)
	check(err)

	// Create new order
	newOrder, err = client.NewCreateOrderService().Symbol(symbol3).
		Side("SELL").Type("MARKET").Quantity(floatQuantity3).
		Do(context.Background())
	check(err)
	fmt.Println(binance_connector.PrettyPrint(newOrder))

	getBalance()

	fmt.Println("End of arbitrage")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	err := godotenv.Load()
	check(err)

	//open pairs.csv
	f, err := os.Open("pairs.csv")
	check(err)

	pairsCSV = f

	defer f.Close()

	// import env vartiables
	var (
		// BINANCE_API_KEY    = os.Getenv("BINANCE_API_KEY")
		// BINANCE_API_SECRET = os.Getenv("BINANCE_API_SECRET")
		BINANCE_API_KEY    = os.Getenv("TEST_BINANCE_API_KEY")
		BINANCE_API_SECRET = os.Getenv("TEST_BINANCE_API_SECRET")
	)
	// The Client can be initiated with apiKey, secretKey and baseURL.
	// The baseURL is optional. If not specified, it will default to "https://api.binance.com".
	client = binance_connector.NewClient(BINANCE_API_KEY, BINANCE_API_SECRET, "https://testnet.binance.vision")
	// Debug Mode
	//client.Debug = true

	go ticker()

	wg.Wait()

}
