package consensus

import (
	"bytes"
)

type Vote struct {
	ValidatorID int
	BlockHash   []byte
	Stake       int
}

func (c *CasperFFG) Vote(validatorID int, blockHash []byte) {
	vote := &Vote{
		ValidatorID: validatorID,
		BlockHash:   blockHash,
		Stake:       c.Validators[validatorID].Stake,
	}
	c.Votes = append(c.Votes, vote)
}

func (c *CasperFFG) CountVotes(blockHash []byte) int {
	totalStake := 0
	for _, vote := range c.Votes {
		if bytes.Equal(vote.BlockHash, blockHash) {
			totalStake += vote.Stake
		}
	}
	return totalStake
}
