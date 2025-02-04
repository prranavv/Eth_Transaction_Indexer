package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

//var interestedAddresses = []string{"0xeBec795c9c8bBD61FFc14A6662944748F299cAcf", "0xDFaa75323fB721e5f29D43859390f62Cc4B600b8"}

var interestedAddy = "0xeBec795c9c8bBD61FFc14A6662944748F299cAcf"

func indexer() {
	cl, err := ethclient.Dial("https://eth-client")
	if err != nil {
		panic(err)
	}
	blockNumber := big.NewInt(int64(21774224))
	block, err := cl.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		panic(err)
	}

	for _, transaction := range block.Transactions() {
		if transaction.To().Hex() == string(interestedAddy) {
			fmt.Println(transaction.Hash())
		}
	}

}

func main() {
	indexer()
}
