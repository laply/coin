package main

import (
	"github.com/laply/coin/cli"
	"github.com/laply/coin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}


