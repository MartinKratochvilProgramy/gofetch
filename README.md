# Fetch stock market data asynchronously

This package utilises parallel capabilities of Google's Go language to fetch stock market data from YahooFinance API.

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

Define your tickers in a string array, ticker names should correspond to the names found on yahoo.finance.com, eg:

```Go
	tickers := []string{"aapl", "msft", "tsla", "amzn", "goog", "nflx", "meta"}

	info := gofetch.FetchAsync(tickers)

	for _, item := range info {
		fmt.Println(item.GetTicker(), item.GetPrice())
	}
```

The above code prints ticker names and prices into console:
```shell
meta 177.16
goog 97.1
aapl 155.33
msft 269.32
amzn 101.16
nflx 361.42
tsla 214.24
```