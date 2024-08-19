package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//connecto to public ehtereum getway

	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatalf("Falid to connect the Ethererum client")
	}

	fmt.Println("Connected to ethereum network")

	var ctx = context.Background()
	account := common.HexToAddress("<Add address here>")
	balance, err := client.BalanceAt(ctx, account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)

	//get
}
