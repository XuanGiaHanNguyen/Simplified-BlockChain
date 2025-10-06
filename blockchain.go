package main

// Use an array to store the order of hash and use an dict to store the hash -> block pair (since simple wont implement dict for now)
type Blockchain struct{
	blocks []*Block // slice of the actual block 
}

// Make it possible to add block into blockchain 
func (bc* Blockchain) AddBlock(data string){
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// initing first block of the chain (genesis block)
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// function to make new blockchain
func NewBlockChain () *Blockchain{
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}