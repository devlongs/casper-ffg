package validator

import (
	"fmt"
	"math/rand"
)

type Validator struct {
	ID       int
	Stake    int
	Honest   bool
	Slashed  bool
	Inactive bool
}

func NewValidator(id, initialStake int, honestyThreshold, inactivityChance float32) *Validator {
	return &Validator{
		ID:       id,
		Stake:    initialStake,
		Honest:   rand.Float32() < honestyThreshold,
		Slashed:  false,
		Inactive: rand.Float32() < inactivityChance,
	}
}

func (v *Validator) Slash(amount int) {
	v.Slashed = true
	v.Stake -= amount
	if v.Stake < 0 {
		v.Stake = 0
	}
	fmt.Printf("Validator %d has been slashed. New stake: %d\n", v.ID, v.Stake)
}

func (v *Validator) Reward(amount int) {
	v.Stake += amount
	fmt.Printf("Validator %d rewarded. New stake: %d\n", v.ID, v.Stake)
}
