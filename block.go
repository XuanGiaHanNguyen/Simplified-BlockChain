package main

import (
	"bytes"
	"encoding/gob"
	"log"
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

// Serialize the block chain
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil { 
		log.Panic(err)
	}

	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err!= nil {
		log.Panic(err)
	}

	return &block
}
