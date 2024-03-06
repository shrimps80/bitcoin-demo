package main

import "fmt"

func main() {
	blockchain := NewBlockchain()
	blockchain.AddBlock("Block 1")
	blockchain.AddBlock("Block 2")
	for i, block := range blockchain.Blocks {

		fmt.Printf("Block %d:\n", i)
		fmt.Printf("Prev. Hash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)

	}

}
