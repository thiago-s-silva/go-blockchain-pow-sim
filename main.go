package main

import (
	"fmt"

	"github.com/thiago-silva/learning-blockchain/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("Sent 1 BTC for John")
	bc.AddBlock("Sent 5 BTC for Smith")
	bc.AddBlock("Sent 5 BTC for Andreas")

	fmt.Println("\n------------------------------------------------------------------")
	for _, block := range bc.Blocks {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Prev Hash %x\n", block.PrevBlockHash)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Println("------------------------------------------------------------------")
	}
}
