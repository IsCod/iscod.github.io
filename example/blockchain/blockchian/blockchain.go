package blockchain

type BlockChain struct {
	Blocks []*Block
}

func NewBlockChain() *BlockChain {
	return &BlockChain{Blocks: make([]*Block, 0)}
}

func (c *BlockChain) AddBlock(prevHash, data []byte) {
	block := NewBlock(prevHash, data)
	work := NewProofWork(block)
	_, h := work.Run()

	block.Hash = h
	c.Blocks = append(c.Blocks, block)
}
