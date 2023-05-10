package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/levigross/grequests"
	"github.com/tidwall/gjson"
	"log"
	"os"
	"time"
)

const uniswapURL = "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	query := `
	{
		pair(id: "0xd3d2e2692501a5c9ca623199d38826e513033a17") {
			id
			token0 {
				id
				symbol
			}
			token1 {
				id
				symbol
			}
			reserve0
			reserve1
			totalSupply
			reserveETH
			reserveUSD
			trackedReserveETH
			token0Price
			token1Price
			volumeToken0
			volumeToken1
			volumeUSD
			untrackedVolumeUSD
			txCount
		}
	}
	`

	data, err := fetchData(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Initial data:", data)

	for range time.Tick(1 * time.Minute) {
		data, err := fetchData(query)
		if err != nil {
			log.Println("Error fetching data:", err)
			continue
		}
		fmt.Println("Updated data:", data)
	}
}

func fetchData(query string) (*gjson.Result, error) {
	apiKey := os.Getenv("GRAPH_API_KEY")
	response, err := grequests.Post(uniswapURL, &grequests.RequestOptions{
		JSON:    map[string]string{"query": query},
		Headers: map[string]string{"Content-Type": "application/json", "Authorization": "Bearer " + apiKey},
	})
	if err != nil {
		return nil, err
	}

	jsonData := gjson.ParseBytes(response.Bytes())
	if !jsonData.Get("errors").Exists() {
		data := jsonData.Get("data")
		return &data, nil
	}

	err = fmt.Errorf("Error from the GraphQL API: %s", jsonData.Get("errors").Raw)
	return nil, err
}
