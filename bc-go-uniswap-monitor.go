package main

import (
	"fmt"
	"time"

	"github.com/levigross/grequests"
	"github.com/tidwall/gjson"
)

const uniswapURL = "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2"
const pairID = "0xd3d2e2692501a5c9ca623199d38826e513033a17" // Replace with your pair ID

func main() {
	for {
		err := fetchData()
		if err != nil {
			fmt.Printf("Error fetching data: %v\n", err)
		}
		time.Sleep(1 * time.Minute)
	}
}

func fetchData() error {
	query := `
	{
		pair(id: "` + pairID + `") {
			id
			token0 {
				symbol
			}
			token1 {
				symbol
			}
			reserve0
			reserve1
			volumeToken0
			volumeToken1
			volumeUSD
		}
	}
	`

	response, err := grequests.Post(uniswapURL, &grequests.RequestOptions{
		JSON:    map[string]string{"query": query},
		Headers: map[string]string{"Content-Type": "application/json"},
	})
	

	if err != nil {
		return err
	}

	data := gjson.GetBytes(response.Bytes(), "data.pair")
	token0 := data.Get("token0.symbol").String()
	token1 := data.Get("token1.symbol").String()
	reserve0 := data.Get("reserve0").Float()
	reserve1 := data.Get("reserve1").Float()
	volumeToken0 := data.Get("volumeToken0").Float()
	volumeToken1 := data.Get("volumeToken1").Float()
	volumeUSD := data.Get("volumeUSD").Float()

	fmt.Printf("Pair: %s-%s\n", token0, token1)
	fmt.Printf("Reserves: %f %s, %f %s\n", reserve0, token0, reserve1, token1)
	fmt.Printf("Trading Volume: %f %s, %f %s, %f USD\n", volumeToken0, token0, volumeToken1, token1, volumeUSD)

	return nil
}
