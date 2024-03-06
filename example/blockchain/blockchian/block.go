package blockchain

import (
	"bytes"
	"strconv"
	"time"
)

type Block struct {
	PrevHash []byte
	Data     []byte
	Hash     []byte
	Time     int64
	Nonce    int64
}

func NewBlock(prevHash, data []byte) *Block {
	return &Block{
		PrevHash: prevHash,
		Data:     data,
		Time:     time.Now().Unix(),
		Hash:     make([]byte, 0),
	}
}

func (block *Block) ToBytes() []byte {
	b := [][]byte{
		block.PrevHash,
		block.Data,
		[]byte(strconv.FormatInt(block.Time, 2)),
		[]byte(strconv.FormatInt(block.Nonce, 2)),
	}
	return bytes.Join(b, []byte{})

}
