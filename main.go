package main

import (
	"fmt"

	"github.com/laply/coin/blockchain"
)

func main() {
	chain := blockchain.GetBlockChain()
	fmt.Println(chain)
}