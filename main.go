package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

var interestedAddresses = map[string]int{"0xeBec795c9c8bBD61FFc14A6662944748F299cAcf": 1, "0xDFaa75323fB721e5f29D43859390f62Cc4B600b8": 1}

func indexer() {
	cl, err := ethclient.Dial("https://eth-client")
	if err != nil {
		panic(err)
	}
	for {
		//blockNumber := big.NewInt(int64(21774224))
		block, err := cl.BlockByNumber(context.Background(), nil)
		if err != nil {
			panic(err)
		}

		for _, transaction := range block.Transactions() {
			_, ok := interestedAddresses[transaction.To().Hex()]
			if ok {
				fmt.Println(transaction.Value())
			}
		}
	}

}

func main() {
	indexer()
}
