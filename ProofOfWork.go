package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"strconv"
)

func IntToHex(num int64) []byte {
    return []byte(strconv.FormatInt(num, 16))
}

type ProofOfWork struct {
	block *Block 
	target *big.Int
}

// Structure that holds a pointer to a block and a pointer to a target:
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))// shifiting left the target 

	pow := &ProofOfWork{b, target}

	return pow 
}

// Preparing the data to hash:
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevDataHash, 
			pow.block.Data, 
			IntToHex(pow.block.TimeStamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		}, 
		[]byte{}, 
	)

	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int // int version of hash 
	var hash [32]byte 
	nonce := 0 // counter 
	maxNonce := math.MaxInt64

	fmt.Printf("Mining the block containing \"%s\"\n ", pow.block.Data)

	// Near inf loop limited by math.MaxInt64
	for nonce < maxNonce {
		data := pow.prepareData(nonce) // Prepare data 
		hash = sha256.Sum256(data) // hash the prepared data 
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:]) // convert hash into a bigger number 

		if hashInt.Cmp(pow.target) == -1 { // compare int with target 
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1 

	return isValid
}