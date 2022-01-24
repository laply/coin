package main

import (
	"fmt"

	"github.com/laply/coin/blockchain"
)

func main() {
	chain := blockchain.GetBlockChain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Forth  Block")


	for _, block := range(chain.AllBlocks()) {
		fmt.Println(*block)
	}

}