package blockchain

import (

	"fmt"
	"sync"

	"github.com/laply/coin/db"
	"github.com/laply/coin/utils"
)

type blockChain struct {
	NewestHash string `json:"newestHash"`
	Height int `json:"height"`
}

var b * blockChain;
var once sync.Once

func (b *blockChain) restore(data []byte){	
	utils.FromByte(b, data)
}

func (b *blockChain) persist() {
	db.SaveBlockChain(utils.ToBytes(b))
}

func (b *blockChain) AddBlock(data string){
  block := createBlock(data, b.NewestHash, b.Height + 1)
  b.NewestHash = block.Hash
  b.Height = block.Height
  b.persist()
} 

func (b *blockChain) Blocks() []*Block{
	var blocks []*Block
	hashCursor := b.NewestHash

	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else { break }
	}
	return blocks
}

func BlockChain() *blockChain {
	if b == nil {
		once.Do(func(){
			b = &blockChain{"", 0}
			fmt.Println("init: ", b.NewestHash, b.Height)
			checkpoint := db.GetCheckpoint()
			if checkpoint == nil {
				b.AddBlock("Genesis Block")
			} else {
				fmt.Println("Restoring...")
				b.restore(checkpoint)
			}
			fmt.Println(b.NewestHash, b.Height)
		})
	}
	
	fmt.Println(b.NewestHash)

	return b
}

