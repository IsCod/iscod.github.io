package blockchain

import (
	"crypto/sha256"
	"fmt"
	"math/big"
)

type ProofWork struct {
	Block  *Block
	target *big.Int
}

func NewProofWork(block *Block) *ProofWork {
	target := big.NewInt(1)
	target = target.Lsh(target, 256-1)
	return &ProofWork{Block: block, target: target}
}

func (p *ProofWork) preData(nonce int64) []byte {
	p.Block.Nonce = nonce
	return p.Block.ToBytes()
}

func (p *ProofWork) Run() (int64, []byte) {
	//target := big.NewInt(1)
	//target = target.Lsh(target, 256-1)

	var nonce int64
	var h = sha256.Sum256(p.preData(nonce))
	for {
		hInt := &big.Int{}
		hInt.SetBytes(h[:])
		if p.target.Cmp(hInt) == 1 {
			fmt.Printf("cmp, %d, %x \n", nonce, h)
			break
		}

		nonce += 1
		h = sha256.Sum256(p.preData(nonce))
	}
	return nonce, h[:]
}
