package blockchain

type BlockChain struct {
	Blocks []*Block // list of block pointers
	Size   int16    // number of blocks in the blockchain
}

// AddBlock appends a new block into the chain
func (bc *BlockChain) AddBlock(data string) {
	// get previous block
	pb := bc.Blocks[len(bc.Blocks)-1]

	// generate new block
	b := NewBlock(data, pb.Hash)

	// add new block into blockchain
	bc.Blocks = append(bc.Blocks, b)
	bc.Size++
}

// NewBlockchain initiates a new blockchain with the genesis block
func NewBlockchain() *BlockChain {
	return &BlockChain{Blocks: []*Block{NewGenesisBlock()}, Size: 1}
}
