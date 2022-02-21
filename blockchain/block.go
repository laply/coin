package blockchain

import (
	"errors"
	"strings"
	"time"

	"github.com/laply/coin/db"
	"github.com/laply/coin/utils"
)

type Block struct {
	Transaction []*Tx 	`json:"transection"`
	Hash string 		`json:"hash"`
	PrevHash string 	`json:"prevHash,omitempty"`
	Height int 			`json:"height"`
	Difficulty int 		`json:"difficulty"`
	Nonce int 			`json:"nonce"`
	TimeStamp int		`json:"timestemp"`
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

func (b *Block) restore(data []byte){
	utils.FromByte(b, data)
}

func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		b.TimeStamp = int(time.Now().Unix())
		hash := utils.Hash(b)
		// fmt.Printf("\n\n\nTarget:%s\nHash:%s\nNonce:%d\n\n\n", target, hash, b.Nonce)

		if strings.HasPrefix(hash, target){
			b.Hash = hash
			break 
		} else {
			b.Nonce++
		}
	}
}

func createBlock(prevHash string, height int) *Block {
	block := &Block {
		Transaction: []*Tx{makeCoinbaseTx("laply")},
		Hash: "",
		PrevHash: prevHash,
		Height: height,
		Difficulty: BlockChain().difficulty(),
		Nonce: 0,
	}

	block.mine()
	block.persist()
	return block
}

var ErrNotFound = errors.New("")

func FindBlock(hash string) (*Block, error){
	blockByte := db.GetBlock(hash)
	if blockByte == nil {
		return nil, ErrNotFound
	}
	block := &Block{}
	block.restore(blockByte)

	return block, nil
}