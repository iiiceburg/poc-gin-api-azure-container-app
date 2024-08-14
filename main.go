package main

import (
    "github.com/gin-gonic/gin"
    "math/rand"
    "net/http"
    "time"
)

type StockPrice struct {
    Symbol string  `json:"symbol"`
    Price  float64 `json:"price"`
}

func generateMockStockPrices(count int) []StockPrice {
    rand.Seed(time.Now().UnixNano())
    symbols := []string{"AAPL", "GOOGL", "MSFT", "AMZN", "TSLA", "META", "NFLX", "NVDA", "ADBE", "INTC", 
	"CSCO", "CRM", "BABA", "ORCL", "IBM", "AMD", "TXN", "QCOM", "WMT", "DIS", 
	"PYPL", "XOM", "BA", "PFE", "V", "MA", "UNH", "KO", "COST", "NKE", 
	"ABT", "MDT", "AVGO", "CSX", "MCD", "HD", "CVX", "GILD", "SBUX", "LOW", 
	"AMGN", "WFC", "T", "PEP", "MS", "DHR", "UNP", "LMT", "HON", "UPS", 
	"NEE", "TMO", "BMY", "NKE", "WBA", "FIS", "MO", "EOG", "C", "CAT", 
	"CL", "GM", "DE", "MDLZ", "TGT", "VZ", "HPE", "USB", "HUM", "SPG", 
	"SYK", "AIG", "BKNG", "BMY", "RCL", "WMT", "MCK", "TRV", "TROW", "SRE", 
	"RBLX", "PLD", "ADP", "EL", "CME", "AON", "MSCI", "CS", "KMB", "WELL", 
	"ITW", "PXD", "CME", "KHC", "QCOM", "HOG", "FISV", "D", "IEX", "PGR", 
	"CB", "ZTS", "SNA", "AZO", "FISV", "BAX", "CERN", "DOW", "PPL", "IT"}
    prices := make([]StockPrice, count)

    for i := 0; i < count; i++ {
        symbol := symbols[rand.Intn(len(symbols))]
        price := 100 + rand.Float64()*900 // Random price between 100 and 1000
        prices[i] = StockPrice{Symbol: symbol, Price: price}
    }

    return prices
}

func main() {
    router := gin.Default()

    router.GET("/stock-prices", func(c *gin.Context) {
        stockPrices := generateMockStockPrices(100)
        c.JSON(http.StatusOK, stockPrices)
    })

    router.Run(":8080")
}
