package blockchain

type block struct {
	data string
	hash string
	prevHash string
}
type blockChain struct {
	blocks []block
}

var b * blockChain;

func GetBlockChain() *blockChain {
	if b == nil {
		b = &blockChain{}
	}

	return b
}

