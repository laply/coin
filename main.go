package main

import (
	"github.com/laply/coin/blockchain"
	"github.com/laply/coin/cli"
	"github.com/laply/coin/db"
)

func main() {
	defer db.Close()

	blockchain.BlockChain()

	cli.Start()
	
}


