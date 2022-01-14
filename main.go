package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data string
	hash string
	prevHash string
}

func main() {
	genesisBlock := block{"genesis Block", "", ""}
	genesisBlock.hash = fmt.Sprintf("%x",sha256.Sum256([]byte(genesisBlock.data+ genesisBlock.prevHash)))
	fmt.Println(genesisBlock)
}