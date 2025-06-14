package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"
)

const targetBits = 23

type ProofOfWork struct {
	block  *Block   // block intendent to be validated
	target *big.Int // the difficulty target
}

// NewProofOfWork represents the proof of work logic
func NewProofOfWork(b *Block) *ProofOfWork {
	t := big.NewInt(1)
	t.Lsh(t, uint(256-targetBits))

	pow := &ProofOfWork{b, t}
	return pow
}

// prepare prepare the data to run the proof of work
func (pow *ProofOfWork) prepare(nonce int) []byte {
	d := bytes.Join([][]byte{
		pow.block.PrevBlockHash,
		pow.block.Data,
		[]byte(strconv.FormatInt(pow.block.Timestamp, 10)),
		[]byte(strconv.FormatInt(int64(targetBits), 10)),
		[]byte(strconv.FormatInt(int64(nonce), 10)),
	},
		[]byte{},
	)

	return d
}

// Run executes the proof of work (mining) until solving the difficulty
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("\n[PoW] Data: '%s'\n", pow.block.Data)
	startTime := time.Now()

	// mining looping (proof of work)
	for nonce < math.MaxInt64 {
		// prepare the proof of work data
		data := pow.prepare(nonce)

		// calculate the hash
		hash = sha256.Sum256(data)
		fmt.Printf("\r[PoW] %x", hash)

		// convert hash to a big number
		hashInt.SetBytes(hash[:])

		// compare the hash with the target
		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf("\n[PoW] Work Time: %s\n", time.Since(startTime))
			break
		} else {
			nonce++
		}
	}

	return nonce, hash[:]
}

// Validate validates if the hash and nonce satifies the proof of work (is valid)
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepare(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.target) == -1
}
