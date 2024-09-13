package blockchain

import (
	"bytes"
)

type Chain struct {
	Blocks []*Block
}

func NewChain() *Chain {
	return &Chain{
		Blocks: []*Block{NewGenesisBlock()},
	}
}

func (c *Chain) AddBlock() *Block {
	parentBlock := c.Blocks[len(c.Blocks)-1]
	newBlock := NewBlock(len(c.Blocks), parentBlock.Hash)
	c.Blocks = append(c.Blocks, newBlock)
	return newBlock
}

func (c *Chain) DetectFork() bool {
	seenParents := make(map[string]bool)
	for _, block := range c.Blocks {
		parentHashStr := string(block.ParentHash)
		if seenParents[parentHashStr] {
			return true
		}
		seenParents[parentHashStr] = true
	}
	return false
}

func (c *Chain) GetBlock(hash []byte) *Block {
	for _, block := range c.Blocks {
		if bytes.Equal(block.Hash, hash) {
			return block
		}
	}
	return nil
}
