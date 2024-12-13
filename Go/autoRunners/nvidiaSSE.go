package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const alphaURL = "https://www.alphavantage.co"
const apiKey = "01EXHL3Z9VT3ZATF"

var etLocation, _ = time.LoadLocation("America/New_York")

// Define a structure to match the JSON response from Alpha Vantage
type GlobalQuote struct {
	Symbol           string `json:"01. symbol"`
	Price            string `json:"05. price"`
	LatestTradingDay string `json:"07. latest trading day"`
}

type CryptoQuote struct {
	Price string `json:"5. Exchange Rate"`
}

// Cache for storing the last known stock price
var cachedPrice string
var lastFetched time.Time

func sseHandler(w http.ResponseWriter, r *http.Request, events <-chan string) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for {
		price := <-events
		fmt.Fprintf(w, "data: %s\n\n", price)
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
	}
}

// fetch function now takes a URL parameter
func fetch(url string) (string, error) {
	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	return string(body), nil
}

func getCurrentCryptoPrice(coinSymbol string) (string, error) {
	url := fmt.Sprintf("%s/query?function=CURRENCY_EXCHANGE_RATE&from_currency=%s&to_currency=USD&apikey=%s", alphaURL, coinSymbol, apiKey)
	responseBody, err := fetch(url)
	if err != nil {
		return "", err
	}

	// Parse the JSON response
	var result map[string]json.RawMessage
	if err := json.Unmarshal([]byte(responseBody), &result); err != nil {
		return "", fmt.Errorf("error parsing JSON: %v", err)
	}

	var quote CryptoQuote
	if err := json.Unmarshal(result["Realtime Currency Exchange Rate"], &quote); err != nil {
		return "", fmt.Errorf("error parsing Realtime Currency Exchange Rate: %v", err)
	}

	// Update cache
	cachedPrice = quote.Price
	lastFetched = time.Now()

	return cachedPrice, nil
}

func getCurrentStockPrice(stockSymbol string) (string, error) {
	url := fmt.Sprintf("%s/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", alphaURL, stockSymbol, apiKey)
	responseBody, err := fetch(url)
	if err != nil {
		return "", err
	}

	// Parse the JSON response
	var result map[string]json.RawMessage
	if err := json.Unmarshal([]byte(responseBody), &result); err != nil {
		return "", fmt.Errorf("error parsing JSON-> %v", err)
	}

	var quote GlobalQuote
	if err := json.Unmarshal(result["Global Quote"], &quote); err != nil {
		return "", fmt.Errorf("error parsing Global Quote-> %v -> %s", err, responseBody)
	}

	// Update cache
	cachedPrice = quote.Price
	lastFetched = time.Now()

	return cachedPrice, nil
}

func isMarketOpen() bool {
	now := time.Now().In(etLocation)

	// Market hours in Eastern Time
	start := time.Date(now.Year(), now.Month(), now.Day(), 9, 30, 0, 0, etLocation)
	end := time.Date(now.Year(), now.Month(), now.Day(), 16, 0, 0, 0, etLocation)

	return now.After(start) && now.Before(end)
}

func stockPriceLoop(stockSymbol string, events chan<- string) {
	for {
		if isMarketOpen() || len(cachedPrice) < 1 {
			price, err := getCurrentStockPrice(stockSymbol)
			if err != nil {
				fmt.Printf("Error fetching stock price-> %v\n", err)
				return
			}

			if price != cachedPrice && len(cachedPrice) > 0 {
				cachedPrice = price
				lastFetched = time.Now()
				events <- price
			}

			time.Sleep(time.Minute) // Adjust frequency as needed
		}
	}
}

func main() {
	stockSymbol := "NVDA" // Nvidia stock symbol
	events := make(chan string)

	go stockPriceLoop(stockSymbol, events)
	go eventHandler(, events)

	http.HandleFunc("/stock-price", func(w http.ResponseWriter, r *http.Request) {
		sseHandler(w, r, events)
	})

	fmt.Println("SSE server is running on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		println(err.Error())
		return
	}
}

//package main
//
//import (
//	"encoding/json"
//	"fmt"
//	"io"
//	"net/http"
//	"time"
//)
//
//const alphaURL = "https://www.alphavantage.co"
//
//const apiKey = "01EXHL3Z9VT3ZATF"
//
//var etLocation, _ = time.LoadLocation("America/New_York")
//
//// Define a structure to match the JSON response from Alpha Vantage
//type GlobalQuote struct {
//	Symbol           string `json:"01. symbol"`
//	Price            string `json:"05. price"`
//	LatestTradingDay string `json:"07. latest trading day"`
//}
//
//type CryptoQuote struct {
//	Price string `json:"5. Exchange Rate"`
//}
//
//// Cache for storing the last known stock price
//var cachedPrice string
//var lastFetched time.Time
//
//func getCurrentCryptoPrice(coinSymbol string) (string, error) {
//	url := fmt.Sprintf("%s/query?function=CURRENCY_EXCHANGE_RATE&from_currency=%s&to_currency=USD&apikey=%s", alphaURL, coinSymbol, apiKey)
//
//	// Make the HTTP GET request
//	resp, err := http.Get(url)
//	if err != nil {
//		return "", fmt.Errorf("error making request: %v", err)
//	}
//	defer resp.Body.Close()
//
//	// Check if the response status is OK
//	if resp.StatusCode != http.StatusOK {
//		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
//	}
//
//	// Read the response body
//	body, err := io.ReadAll(resp.Body)
//	if err != nil {
//		return "", fmt.Errorf("error reading response body: %v", err)
//	}
//
//	// Parse the JSON response
//	var result map[string]json.RawMessage
//	if err := json.Unmarshal(body, &result); err != nil {
//		return "", fmt.Errorf("error parsing JSON: %v", err)
//	}
//
//	if err := json.Unmarshal(body, &result); err != nil {
//		return "", fmt.Errorf("error parsing JSON: %v", err)
//	}
//
//	var quote CryptoQuote
//	if err := json.Unmarshal(result["Realtime Currency Exchange Rate"], &quote); err != nil {
//		return "", fmt.Errorf("error parsing Realtime Currency Exchange Rate: %v", err)
//	}
//
//	// Update cache
//	cachedPrice = quote.Price
//	lastFetched = time.Now()
//
//	return cachedPrice, nil
//}
//
//func fetch(stockSymbol string) (string, error) {
//	url := fmt.Sprintf("%s/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", alphaURL, stockSymbol, apiKey)
//
//	// Make the HTTP GET request
//	resp, err := http.Get(url)
//	if err != nil {
//		return "", fmt.Errorf("error making request: %v", err)
//	}
//	defer resp.Body.Close()
//
//	// Check if the response status is OK
//	if resp.StatusCode != http.StatusOK {
//		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
//	}
//
//	// Read the response body
//	body, err := io.ReadAll(resp.Body)
//	if err != nil {
//		return "", fmt.Errorf("error reading response body: %v", err)
//	}
//
//	// Parse the JSON response
//	var result map[string]json.RawMessage
//	if err := json.Unmarshal(body, &result); err != nil {
//		return "", fmt.Errorf("error parsing JSON: %v", err)
//	}
//
//	// Print the raw JSON response for debugging
//	fmt.Printf("Raw API response: %s\n", body)
//
//	var quote GlobalQuote
//	if err := json.Unmarshal(result["Global Quote"], &quote); err != nil {
//		return "", fmt.Errorf("error parsing Global Quote: %v", err)
//	}
//
//	return quote.Price, nil
//}
//
//func getCurrentStockPrice(stockSymbol string) (string, error) {
//	price, err := fetch(stockSymbol)
//	if err != nil {
//		return "", err
//	}
//
//	// Update cache
//	cachedPrice = price
//	lastFetched = time.Now()
//
//	return cachedPrice, nil
//}
//
//func isMarketOpen() bool {
//	now := time.Now().In(time.FixedZone("ET", -5*3600)) // Eastern Time (ET)
//
//	// Market hours in Eastern Time
//	start := time.Date(now.Year(), now.Month(), now.Day(), 9, 30, 0, 0, etLocation)
//	end := time.Date(now.Year(), now.Month(), now.Day(), 16, 0, 0, 0, etLocation)
//
//	return now.After(start) && now.Before(end)
//}
//
//func main() {
//	//price, err := getCurrentCryptoPRice("")
//	//if err != nil {
//	//	return
//	//}
//	//println(price)
//
//	stock := "nvda"
//	var price string
//	var err error
//
//	for {
//		if isMarketOpen() || len(price) < 1 {
//			price, err = getCurrentStockPrice(stock)
//			if err != nil {
//				fmt.Printf("Error fetching stock price: %v\n", err)
//			} else {
//				fmt.Printf("%s current stock price: $%s\n", stock, price)
//			}
//		} else {
//			if cachedPrice == "" {
//				fmt.Println("Market is closed, and no cached data is available.")
//			} else {
//				fmt.Printf("%s current stock price (cached): $%s\n", stock, cachedPrice)
//			}
//		}
//
//		//Request limit 25 x per
//		time.Sleep(time.Hour)
//	}
//}
