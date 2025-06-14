package blockchain

import (
	"time"
)

type Block struct {
	Timestamp     int64  // when the block was created
	Data          []byte // payload
	PrevBlockHash []byte // previous block hash
	Hash          []byte // current block hash (PrevBlockHash + Data + Timestamp)
	Nonce         int    // used to handle the PoW
}

// NewBlock initiates a new block based on the data and the previous block hash
func NewBlock(data string, prevBlockHash []byte) *Block {
	// init block
	b := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
	}

	// initiates a new proof of work component
	pow := NewProofOfWork(b)

	// executes the proof of work until finding a hash that satisfies the target based on the difficulty
	nonce, hash := pow.Run()

	// set the found values
	b.Nonce = nonce
	b.Hash = hash[:]

	return b
}

// NewGenesisBlock initiates the first genesis block with empty data and hash
func NewGenesisBlock() *Block {
	return NewBlock("genesis", []byte{})
}
