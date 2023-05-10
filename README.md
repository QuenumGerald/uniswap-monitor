# Uniswap Monitoring Program

This Go program queries the Uniswap V2 subgraph using the GraphQL API to fetch the pair data, including liquidity, trading volume, and the number of users. The program runs every minute to provide updated information.

## Prerequisites

- Go programming language installed
- An API key for The Graph's GraphQL API

## Dependencies

- github.com/joho/godotenv
- github.com/levigross/grequests
- github.com/tidwall/gjson

## Setup

### 1. Create the project directory

Create a new directory for the project and navigate to it:

mkdir uniswap-monitor
cd uniswap-monitor


### 2. Environment file

Create a `.env` file in the project directory with the following content:

Replace `your_api_key_here` with your actual API key.

### 3. Install Go libraries

Install the required Go libraries:

go get github.com/joho/godotenv
go get github.com/levigross/grequests
go get github.com/tidwall/gjson


### 4. Create the Go source file

Create a new file called `bc-go-uniswap-monitor.go` in the project directory and add the previously provided Go code to it.

## Usage

### Running the program

Run the program using the following command:


The program will output the fetched pair data every minute, including liquidity, trading volume, and number of users.
