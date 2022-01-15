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

type blockChain struct {
	blocks []block
}
func (bc *blockChain) makeHash(data string) string {
	return fmt.Sprintf("%x",sha256.Sum256([]byte(data)))
}

func (bc *blockChain) getLastHash() string {
	if(len(bc.blocks) != 0) {
		return  bc.blocks[len(bc.blocks)-1].hash
	}
	return ""
}

func (bc *blockChain) addBlock(data string){
	bc.blocks = append(bc.blocks, block{data, bc.makeHash(data),  bc.getLastHash()})
}

func (bc blockChain) listBlocks(){
	for _, data := range bc.blocks {
		fmt.Printf("Data: %s\nHash: %s\nPrev Hash: %s\n", data.data, data.hash, data.prevHash)
	}
}


func main() {
	chain := blockChain{}
	chain.addBlock("1st")
	chain.addBlock("2nd")
	chain.addBlock("3rd")
	chain.listBlocks()
}