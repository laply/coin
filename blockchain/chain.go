package blockchain

import (
	"fmt"
	"sync"

	"github.com/laply/coin/db"
	"github.com/laply/coin/utils"
)
const (
	defaultDifficulty 	int = 2
	difficultyInterval 	int = 5
	blockInterval 	   	int = 2
	rangeTime 			int = 3

)

type blockChain struct {
	NewestHash string `json:"newestHash"`
	Height int `json:"height"`
	CurrentDifficulty int `json:"currentDifficulty"`
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
  b.CurrentDifficulty = block.Difficulty
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

func (b *blockChain) reCalculateDifficulty() int {
	
	allBlocks := b.Blocks()
	newestBlock := allBlocks[0]
	lastRecalculatedBlock := allBlocks[difficultyInterval-1]

	actualTime := newestBlock.TimeStamp/60 - lastRecalculatedBlock.TimeStamp/60
	expectedTime := difficultyInterval*blockInterval

	if actualTime < (expectedTime - rangeTime) {
		return b.CurrentDifficulty + 1
	} else if actualTime > (expectedTime + rangeTime) {
		return b.CurrentDifficulty - 1 
	} 
	return b.CurrentDifficulty
}


func (b *blockChain) difficulty() int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height % difficultyInterval == 0 {
		return b.reCalculateDifficulty()
	} else {
		return b.CurrentDifficulty
	}
}

func BlockChain() *blockChain {
	if b == nil {
		once.Do(func(){
			b = &blockChain{
				Height: 0,
			}
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
	return b
}

