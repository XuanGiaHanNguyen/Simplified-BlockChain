package main

import (
	"log"
	bolt "go.etcd.io/bbolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

// Use an array to store the order of hash and use an dict to store the hash -> block pair (since simple wont implement dict for now)
type Blockchain struct{
	tip []byte
	db *bolt.DB 
}

type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

// Make it possible to add block into blockchain 
func (bc* Blockchain) AddBlock(data string){
	var lastHash []byte 
	
	err := bc.db.View(func(tx *bolt.Tx) error {
		b:= tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil 
	})

	if err != nil {
		log.Panic(err)
	}

	NewBlock := NewBlock(data, lastHash)
	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(NewBlock.Hash, NewBlock.Serialize())

		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), NewBlock.Hash)
		bc.tip = NewBlock.Hash

		return nil 
	})
}

// initing first block of the chain (genesis block)
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// function to make new blockchain
func NewBlockChain () *Blockchain{
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)

	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b:= tx.Bucket([]byte(blocksBucket))
		if b == nil{
			genesis := NewGenesisBlock()
			b,err := tx.CreateBucket([]byte(blocksBucket))

			if err != nil {
				log.Panic(err)
			}

			err = b.Put(genesis.Hash, genesis.Serialize())
			err = b.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})
	bc := Blockchain{tip, db}

	return &bc
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}

	return bci
}

func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	i.currentHash = block.PrevDataHash

	return block
}