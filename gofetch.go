package gofetch

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/tidwall/gjson"
)

func fetchPrevClose(ticker string) float64 {
	resp, err := http.Get("https://query1.finance.yahoo.com/v8/finance/chart/" + ticker)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	res_1 := strings.Replace(sb, "[", "", -1)
	res_2 := strings.Replace(res_1, "[", "", -1)
	result := gjson.Get(res_2, "chart.result.meta.chartPreviousClose").Float()

	return result
}

type TickerInfo struct {
	ticker string
	price  float64
}

func FetchAsync(tickers []string) []TickerInfo {
	var wg sync.WaitGroup
	wg.Add(len(tickers))

	var res []TickerInfo

	for i, ticker := range tickers {
		go func(i int, ticker string) {
			price := fetchPrevClose(ticker)
			res = append(res, TickerInfo{ticker: ticker, price: price})
			wg.Done()
		}(i, ticker)
	}
	wg.Wait()

	return res
}
