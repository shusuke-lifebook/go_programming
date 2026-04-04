package main

import (
	"fmt"

	"github.com/markcheno/go-quote"
	a "github.com/markcheno/go-talib"
)

func main() {
	// Coinbase からデータ取得（BTC-USD の例）
	spy, err := quote.NewQuoteFromCoinbase("BTC-USD", "2021-04-01", "2021-04-04", quote.Daily)
	if err != nil {
		panic(err)
	}

	fmt.Print(spy.CSV())

	// RSI 計算
	rsi2 := a.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
}
