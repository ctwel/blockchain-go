package main

import (
	"crypto/sha256"
	"fmt"
)

// Block 0. 定义结构
type Block struct {
	// 1. 前区块hash
	PrevHash []byte
	// 2. 当前区块hash
	Hash []byte
	// 3. 数据
	Data []byte
}

// NewBlock 2. 创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		PrevHash: prevBlockHash,
		Hash:     []byte{}, // 先填空，后面在计算 //TODO
		Data:     []byte(data),
	}
	block.SetHash()
	return &block
}

// SetHash 3. 生成Hash
func (block *Block) SetHash() {
	// 1. 拼装数据
	blockInfo := append(block.PrevHash, block.Data...)
	// 2. sha256
	// func Sum256(data []byte) [Size]byte {
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}

// BlockChain 4. 定义区块链结构
type BlockChain struct {
	// 定义一个区块链数组
	blocks []*Block
}

// NewBlockChain 5. 定义一个区块链
func NewBlockChain() *BlockChain {
	// 创建一个创世块，并作为第一个区块添加到区块链中
	gensisBlock := GensisBlock()
	return &BlockChain{
		blocks: []*Block{
			gensisBlock,
		},
	}
}

// GensisBlock 定义一个创世块
func GensisBlock() *Block {
	return NewBlock("创世块", []byte{})
}

// AddBlock 6. 添加区块
func (bc *BlockChain) AddBlock(data string) {
	lastBlock := bc.blocks[len(bc.blocks)-1]
	prevHash := lastBlock.Hash
	block := NewBlock(data, prevHash)
	bc.blocks = append(bc.blocks, block)
}

func main() {
	bc := NewBlockChain()
	bc.AddBlock("班长向班花转了50枚比特币")
	bc.AddBlock("班长又向班花转了50枚比特币")
	for i, block := range bc.blocks {
		fmt.Printf("====== 当前区块高度：%d =======\n", i)
		fmt.Printf("前区块哈希值：%x\n", block.PrevHash)
		fmt.Printf("当前区块哈希值：%x\n", block.Hash)
		fmt.Printf("当前区块数据：%s\n\n", block.Data)
	}
}
