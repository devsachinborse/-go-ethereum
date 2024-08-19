package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Connect to Cloudflare's public Ethereum gateway
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	defer client.Close()

	// Print confirmation of the connection
	fmt.Println("Connected to Ethereum network")

	// Get the latest block number
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatalf("Failed to get the latest block number: %v", err)
	}

	fmt.Printf("Latest block number: %d\n", blockNumber)
}
