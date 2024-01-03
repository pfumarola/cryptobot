package internal

import (
	"context"

	binance_connector "github.com/binance/binance-connector-go"
)

type ExchangeInfo struct {
	Symbols []*binance_connector.SymbolInfo
}

func (e *ExchangeInfo) NewExchangeInfo(client *binance_connector.Client) *ExchangeInfo {
	exchangeInfo, err := client.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		panic(err)
	}
	e.Symbols = exchangeInfo.Symbols
	return e
}

func (e *ExchangeInfo) GetSymbolLotSize(symbol string) *binance_connector.SymbolFilter {
	for _, s := range e.Symbols {
		if s.Symbol == symbol {
			return s.Filters[1]
		}
	}
	return nil
}
