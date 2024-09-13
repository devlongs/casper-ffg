package test

import (
	"testing"

	"github.com/devlongs/casper-ffg/config"
	"github.com/devlongs/casper-ffg/consensus"
)

func TestSimulation(t *testing.T) {
	cfg := config.DefaultConfig()
	casper, err := consensus.NewCasperFFG(cfg)
	if err != nil {
		t.Fatalf("Failed to create CasperFFG: %v", err)
	}

	initialChainLength := len(casper.Chain.Blocks)

	casper.SimulateRound()

	if len(casper.Chain.Blocks) != initialChainLength+1 {
		t.Errorf("Expected chain length to increase by 1, got %d", len(casper.Chain.Blocks))
	}

	if len(casper.Votes) == 0 {
		t.Error("Expected votes to be cast, but no votes were recorded")
	}
}
