package main

import (
	"fmt"
	"strconv"
	"github.com/tkivite/myGoBlockchain/blockchain"
)


func main(){
	chain := blockchain.InitBlockChain()
	chain.AddBlock("first block after genesis")
	chain.AddBlock("second block after genesis")
	chain.AddBlock("third block after genesis")
	for _, block := range chain.Blocks {
        fmt.Printf("Previous hash: %x\n", block.PrevHash)
        fmt.Printf("data: %s\n", block.Data)
        fmt.Printf("hash: %x\n", block.Hash)

		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
        fmt.Println()

    }
	
}

// func (b *Block) DeriveHash() {	
// 	// This will join our previous block's relevant info with the new blocks
// 	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
// 	//This performs the actual hashing algorithm
// 	hash := sha256.Sum256(info)
// 	b.Hash = hash[:]
// }