package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// "reflect"
	// "strconv"
	// "time"
	"io/ioutil"
	"os"
)

/*
ex:
https://api.coinmarketcap.com/v1/ticker/bitcoin/
*/

// Coin ...
// struct for the json of a coin
type Coin []struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Rank             string `json:"rank"`
	PriceUsd         string `json:"price_usd"`
	PriceBtc         string `json:"price_btc"`
	Two4HVolumeUsd   string `json:"24h_volume_usd"`
	MarketCapUsd     string `json:"market_cap_usd"`
	AvailableSupply  string `json:"available_supply"`
	TotalSupply      string `json:"total_supply"`
	PercentChange1H  string `json:"percent_change_1h"`
	PercentChange24H string `json:"percent_change_24h"`
	PercentChange7D  string `json:"percent_change_7d"`
	LastUpdated      string `json:"last_updated"`
}

// Hist ...
type Hist []struct {
	data        []string `json:"history"`
	LastUpdated string   `json:"last_updated"`
}

func getCurrent(c Coin) error {
	file, e := ioutil.ReadFile("btc.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	err := json.Unmarshal(file, &c)
	if err != nil {
		return err
	}
	return nil
}

// CallAPI ...
// takes a coin name ang called the coinmarketcap api
// for that coin.
// Returns the unparsed json response
func CallAPI(s string) []byte {
	/* coin := []string{"bitcoin", "ethereum", "dash", */
	// "litecoin", "Monero", "MaidSafeCoin",
	// "Stellar", "BitShares", "Dogecoin",
	// "Clams",
	/* } */

	url := fmt.Sprintf("https://api.coinmarketcap.com/v1/ticker/%s", s)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}
	defer resp.Body.Close()

	body, er := ioutil.ReadAll(resp.Body)
	if er != nil {
		panic(er)
	}

	return []byte(body)
}
