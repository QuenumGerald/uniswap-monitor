package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GraphQLRequest struct {
	Query string `json:"query"`
}

type GraphQLResponse struct {
	Data   map[string]interface{} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

func main() {
	// Replace with the appropriate Uniswap subgraph URL (v2 or v3)
	uniswapURL := "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2"

	// Replace with the specific pair's ID
	pairID := "0x00004ee988665cdda9a1080d5792cecd16dc1220"

	query := fmt.Sprintf(`
		{
			pair(id: "%s") {
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
			}
		}
	`, pairID)

	request := GraphQLRequest{Query: query}
	response, err := makeGraphQLRequest(uniswapURL, request)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(response.Errors) > 0 {
		fmt.Println("GraphQL errors:", response.Errors)
		return
	}

	pairData := response.Data["pair"].(map[string]interface{})
	token0Symbol := pairData["token0"].(map[string]interface{})["symbol"].(string)
	token1Symbol := pairData["token1"].(map[string]interface{})["symbol"].(string)
	liquidityToken0 := pairData["reserve0"].(string)
	liquidityToken1 := pairData["reserve1"].(string)
	volumeToken0 := pairData["volumeToken0"].(string)
	volumeToken1 := pairData["volumeToken1"].(string)

	fmt.Printf("Pair: %s-%s\n", token0Symbol, token1Symbol)
	fmt.Printf("Liquidity: %s %s, %s %s\n", liquidityToken0, token0Symbol, liquidityToken1, token1Symbol)
	fmt.Printf("Trading Volume: %s %s, %s %s\n", volumeToken0, token0Symbol, volumeToken1, token1Symbol)
}

func makeGraphQLRequest(url string, request GraphQLRequest) (*GraphQLResponse, error) {
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response GraphQLResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
