package main

import (
	"fmt"
	"strconv"
)

func main() {
	bc := NewBlockChain()

	bc.AddBlock("Testing Blockchain 1")
	bc.AddBlock("Testing Blockchain 2")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevDataHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}