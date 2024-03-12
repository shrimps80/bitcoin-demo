package main

import (
	"fmt"
	"strconv"
)

func main() {
	blockchain := NewBlockchain()
	blockchain.AddBlock("Block 1")
	blockchain.AddBlock("Block 2")
	for i, block := range blockchain.Blocks {

		fmt.Printf("Block %d:\n", i)
		fmt.Printf("Prev. Hash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)

		pow := NewProofOfWork(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.IsValid()))

	}

}
