package config

type Config struct {
	NumValidators    int
	InitialStake     int
	HonestyThreshold float32
	InactivityChance float32
	RewardAmount     int
	SlashAmount      int
}

func DefaultConfig() *Config {
	return &Config{
		NumValidators:    10,
		InitialStake:     100,
		HonestyThreshold: 0.7,
		InactivityChance: 0.1,
		RewardAmount:     10,
		SlashAmount:      50,
	}
}
