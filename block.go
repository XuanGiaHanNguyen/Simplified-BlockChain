package main

import (
	"time"
)


const targetBits = 24 // Stores difficulty at which the block was minded 

// Target adjusting algorithm 

type Block struct { 
	TimeStamp int64
	Data []byte 
	PrevDataHash []byte
	Hash []byte
	Nonce int 
}

// creating a new block 
func NewBlock(data string, prevBlockHash []byte) *Block{
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
