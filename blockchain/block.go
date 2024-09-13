package blockchain

import (
	"time"

	"github.com/devlongs/casperffg/utils"
)

type Block struct {
	Height     int
	Hash       []byte
	ParentHash []byte
	Timestamp  time.Time
}

func NewGenesisBlock() *Block {
	genesisBlock := &Block{
		Height:     0,
		ParentHash: make([]byte, 32),
		Timestamp:  time.Now(),
	}
	genesisBlock.Hash = utils.CalculateHash(genesisBlock.Bytes())
	return genesisBlock
}

func NewBlock(height int, parentHash []byte) *Block {
	block := &Block{
		Height:     height,
		ParentHash: parentHash,
		Timestamp:  time.Now(),
	}
	block.Hash = utils.CalculateHash(block.Bytes())
	return block
}

func (b *Block) Bytes() []byte {
	return append(b.ParentHash, append(utils.IntToBytes(b.Height), b.Timestamp.UnixNano())...)
}
