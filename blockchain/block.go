package blockchain
import (
	"bytes"
	"encoding/gob"
	"log"
)
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int 
}

// type BlockChain struct {
// 	Blocks []*Block
// }

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	// block.DeriveHash()
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

// func (chain *BlockChain) AddBlock(data string) {
// 	prevBlock := chain.Blocks[len(chain.Blocks)-1]
// 	new := CreateBlock(data, prevBlock.Hash)
// 	chain.Blocks = append(chain.Blocks,new)
// }

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// func InitBlockChain() *BlockChain {
// 	return &BlockChain{[]*Block{Genesis()}}
// }

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(b)
	Handle(err)
	return res.Bytes()
}

func Deserialize(data []byte) *Block {
    var block Block
    decoder := gob.NewDecoder(bytes.NewReader(data))
    err := decoder.Decode(&block)
    Handle(err)
    return &block
}