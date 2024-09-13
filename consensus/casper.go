package consensus

import (
	"fmt"
	"math/rand"

	"github.com/devlongs/casper-ffg/blockchain"
	"github.com/devlongs/casper-ffg/config"
	"github.com/devlongs/casper-ffg/validator"
)

type CasperFFG struct {
	Validators     []*validator.Validator
	Chain          *blockchain.Chain
	Votes          []*Vote
	FinalizedBlock *blockchain.Block
	ForkExists     bool
	Config         *config.Config
}

func NewCasperFFG(cfg *config.Config) (*CasperFFG, error) {
	if cfg.NumValidators <= 0 {
		return nil, fmt.Errorf("number of validators must be positive")
	}

	validators := make([]*validator.Validator, cfg.NumValidators)
	for i := 0; i < cfg.NumValidators; i++ {
		validators[i] = validator.NewValidator(i, cfg.InitialStake, cfg.HonestyThreshold, cfg.InactivityChance)
	}

	return &CasperFFG{
		Validators: validators,
		Chain:      blockchain.NewChain(),
		Config:     cfg,
	}, nil
}

func (c *CasperFFG) SimulateRound() {
	newBlock := c.Chain.AddBlock()

	for _, v := range c.Validators {
		if v.Slashed || v.Inactive {
			continue
		}

		if v.Honest {
			c.Vote(v.ID, newBlock.Hash)
		} else {
			randomBlockIndex := rand.Intn(len(c.Chain.Blocks))
			c.Vote(v.ID, c.Chain.Blocks[randomBlockIndex].Hash)
		}
	}

	for _, v := range c.Validators {
		if v.Slashed {
			continue
		}
		if !v.Honest || v.Inactive {
			v.Slash(c.Config.SlashAmount)
		} else {
			v.Reward(c.Config.RewardAmount)
		}
	}

	c.Finalize()

	if c.Chain.DetectFork() {
		fmt.Println("Fork detected in the chain!")
		c.ForkExists = true
	}
}

func (c *CasperFFG) Finalize() {
	totalStake := 0
	for _, v := range c.Validators {
		if !v.Slashed {
			totalStake += v.Stake
		}
	}
	threshold := totalStake * 2 / 3

	for i := len(c.Chain.Blocks) - 1; i >= 0; i-- {
		votes := c.CountVotes(c.Chain.Blocks[i].Hash)
		if votes >= threshold {
			c.FinalizedBlock = c.Chain.Blocks[i]
			break
		}
	}
}

func (c *CasperFFG) PrintStats() {
	fmt.Printf("Total blocks: %d\n", len(c.Chain.Blocks))
	fmt.Printf("Total votes: %d\n", len(c.Votes))
	if c.FinalizedBlock != nil {
		fmt.Printf("Finalized block height: %d\n", c.FinalizedBlock.Height)
	} else {
		fmt.Println("No block finalized yet")
	}

	honestCount := 0
	slashedCount := 0
	totalStake := 0
	for _, v := range c.Validators {
		if v.Honest {
			honestCount++
		}
		if v.Slashed {
			slashedCount++
		}
		totalStake += v.Stake
	}

	fmt.Printf("Honest validators: %d/%d\n", honestCount, len(c.Validators))
	fmt.Printf("Slashed validators: %d/%d\n", slashedCount, len(c.Validators))
	fmt.Printf("Total stake: %d\n", totalStake)
}
