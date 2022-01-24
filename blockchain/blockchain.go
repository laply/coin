package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	data string
	hash string
	prevHash string
}
type blockChain struct {
	blocks []*block
}

var b * blockChain;
var once sync.Once

func createBlock (data string) * block {
	newBlock := block{data, "", getPrevHash()}
	newBlock.calculateHash()
	return &newBlock
}

func (b *block) calculateHash () {
	hash := sha256.Sum256([]byte(b.data + b.prevHash))
	b.hash = fmt.Sprintf("%x", hash)
}

func getPrevHash() string {
	totalBlock := len(GetBlockChain().blocks)
	if totalBlock == 0 { return "" } 
	return GetBlockChain().blocks[totalBlock-1].hash
}


func (b *blockChain) AddBlock(data string){
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockChain() *blockChain {
	if b == nil {
		once.Do(func(){
			b = &blockChain{}
			b.AddBlock("Genesis Block")
		})
	}

	return b
}

func (b *blockChain) AllBlocks() []*block{
	return b.blocks
}