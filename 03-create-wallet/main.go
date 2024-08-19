package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// Create private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// Convert private key to a readable format
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("Private Key: ", hexutil.Encode(privateKeyBytes)[2:]) // Strip off the 0x

	// Generate public key from private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey) // Convert to *ecdsa.PublicKey
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)           // Convert to bytes
	fmt.Println("Public Key: ", hexutil.Encode(publicKeyBytes)[4:]) // Strip off the 0x and the first 2 characters

	// Generate address
	// Keccak-256 hash of the public key, then take the last 20 bytes (40 characters) and prefix it with 0x
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("Address: ", address)
}
