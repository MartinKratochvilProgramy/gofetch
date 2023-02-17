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
