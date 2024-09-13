package blockchain

import (
	"encoding/binary"
	"time"

	"github.com/devlongs/casper-ffg/utils"
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
	heightBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(heightBytes, uint32(b.Height))

	timestampBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(timestampBytes, uint64(b.Timestamp.UnixNano()))

	return append(b.ParentHash, append(heightBytes, timestampBytes...)...)
}
