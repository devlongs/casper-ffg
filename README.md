# Casper FFG Simulation

This project implements a simplified version of the Casper FFG (Friendly Finality Gadget) consensus mechanism in Go. It simulates a blockchain network with validators, voting, and finalization processes.

## Project Structure

```
casperffg/
│
├── cmd/
│   └── simulator/
│       └── main.go
├── blockchain/
│   ├── block.go
│   └── chain.go
├── consensus/
│   ├── casper.go
│   └── vote.go
├── validator/
│   └── validator.go
├── config/
│   └── config.go
├── utils/
│   └── hashing.go
├── test/
│   └── simulation_test.go
├── README.md
└── go.mod
```

## How to Run

1. Clone the repository
2. Navigate to the project directory
3. Run `go mod tidy` to ensure all dependencies are correctly managed
4. Run the simulation with `go run cmd/simulator/main.go`

## Running Tests

To run the tests, use the following command:

```
go test ./test
```

## Future Improvements

- Implement more sophisticated fork resolution mechanisms
- Add support for transactions within blocks
- Implement a simple API for interacting with the simulation
- Improve test coverage and add benchmarks
- Add logging and error handling