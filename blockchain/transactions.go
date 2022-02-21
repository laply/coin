package blockchain

import (
	"time"

	"github.com/laply/coin/utils"
)


const (
	minerReward int = 50
)
type Tx struct{
	Id string 		`json:"id"`
	Timestamp int 	`json:"timestamp"`
	TxIns []*TxIn 	`json:"txIns"`
	TxOuts []*TxOut `json:"txOut"`
}

func (t *Tx) getId(){
	t.Id = utils.Hash(t)
}

type TxIn struct{
	Owner string 	`json:"owner"`
	Amount int 		`json:"amount"`
}

type TxOut struct{
	Owner string 	`json:"owner"`
	Amount int		`json:"amount"`
}

func makeCoinbaseTx(address string) *Tx{

	tx := Tx{
		Id: "",
		Timestamp: int(time.Now().UnixNano()),
		TxIns: []*TxIn{{"COINBASE", minerReward}},
		TxOuts: []*TxOut{{address, minerReward}},
	}

	tx.getId()
	return &tx
}