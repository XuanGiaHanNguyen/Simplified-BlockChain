package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct { 
	TimeStamp int64
	Data []byte 
	PrevDataHash []byte
	Hash []byte
}

// creating the hash pass using the provided info 
func (b *Block) setHash(){
	timestamp := []byte(strconv.FormatInt(b.TimeStamp, 10))
	headers := bytes.Join([][]byte{b.PrevDataHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// creating a new block 
func NewBlock(data string, prevBlockHash []byte) *Block{
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.setHash()
	return block
}