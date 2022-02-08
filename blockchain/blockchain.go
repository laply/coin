package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
)

type Block struct {
	Data string `json:"data"`
	Hash string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height int `json:"height"`
}
type blockChain struct {
	blocks []*Block
}

var b * blockChain;
var once sync.Once

func createBlock (data string) * Block {
	newBlock := Block{data, "", getPrevHash(),len(GetBlockChain().blocks)+1}
	newBlock.calculateHash()
	return &newBlock
}

func (b *Block) calculateHash () {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func getPrevHash() string {
	totalBlock := len(GetBlockChain().blocks)
	if totalBlock == 0 { return "" } 
	return GetBlockChain().blocks[totalBlock-1].Hash
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

func (b *blockChain) AllBlocks() []*Block {
	return b.blocks
}

var ErrNotFound = errors.New("block not found")

func (b *blockChain) FindBlock(height int) (*Block, error) {
	if height > len(b.blocks){
		return nil, ErrNotFound
	}

	return b.blocks[height - 1], nil
}