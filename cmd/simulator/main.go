package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/devlongs/casperffg/config"
	"github.com/devlongs/casperffg/consensus"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	cfg := config.DefaultConfig()

	casper, err := consensus.NewCasperFFG(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize CasperFFG: %v", err)
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("===== Round %d =====\n", i+1)
		casper.SimulateRound()
		if casper.FinalizedBlock != nil {
			fmt.Printf("Finalized block height: %d\n", casper.FinalizedBlock.Height)
		} else {
			fmt.Println("No block finalized yet")
		}
		if casper.ForkExists {
			fmt.Println("Warning: Fork detected in the chain!")
		}
		fmt.Println()
	}

	casper.PrintStats()
}
