# Fetch stock market data asynchronously

This package utilises concurrency capabilities of Google's Go language to fetch stock market data from YahooFinance API.

When fetching 7 stock tickers ("aapl", "msft", "tsla", "amzn", "goog", "nflx", "meta") there was approximatedly 2.66x speedup, actual speedup will vary on the parallel capabilities of your machine.

## Installation

```bash
$ go get -u github.com/MartinKratochvilProgramy/gofetch
```

## Quick Start

Add this import line to the file you're working in:

```Go
import "github.com/MartinKratochvilProgramy/gofetch"
```

Define your tickers in a string array, ticker names should correspond to the names found on yahoo.finance.com. Function gofetch.FetchAsync() returns array of structs TickerInfo{ticker string, price float64}. To get the ticker name or price, use the GetTicker() or GetPrice() functions.

```Go
tickers := []string{"aapl", "msft", "tsla", "amzn", "goog", "nflx", "meta"}

info := gofetch.FetchAsync(tickers)

for _, item := range info {
    fmt.Println(item.GetTicker(), item.GetPrice())
}
```

The above code prints ticker names and previous day close prices into console:
```shell
meta 177.16
goog 97.1
aapl 155.33
msft 269.32
amzn 101.16
nflx 361.42
tsla 214.24
```

## Using the FetchPrevClose function

If you do not wish to load ticker information into an array of structs, use the gofetch.FetchPrevClose() which returns previous day close price for individual stock.

```Go
info := gofetch.FetchPrevClose("aapl")
fmt.Println(info.GetTicker(), info.GetPrice())
```

Returns:
```shell
aapl 155.33
```
