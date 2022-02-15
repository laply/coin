package blockchain

import (
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

func (b *blockChain) persist() {
	db.SaveBlockChain(utils.ToBytes(b))
}

func (b *blockChain) AddBlock(data string){
  block := createBlock(data, b.NewestHash, b.Height + 1)
  b.NewestHash = block.Hash
  b.Height = block.Height
  b.persist()

} 

func BlockChain() *blockChain {
	if b == nil {
		once.Do(func(){
			b = &blockChain{"", 0}
			b.AddBlock("Genesis Block")
		})
	}

	return b
}

