package main

import "github.com/iscod/example/blockchain/blockchain"
import "fmt"

func main() {
	fmt.Println("Starting")
	chain := blockchain.NewBlockChain()
	chain.AddBlock(make([]byte, 32), []byte("fist to gold"))

}
