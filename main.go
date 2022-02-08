package main

import (
	"github.com/laply/coin/explorer"
	"github.com/laply/coin/rest"
)

func main() {
	go explorer.Start(4200)
	rest.Start(4210)
}


