package main

import "fmt"

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
	return &block
}

func main() {
	block := NewBlock("老师转班长一枚比特币", []byte{})
	fmt.Printf("前区块哈希值：%x\n", block.PrevHash)
	fmt.Printf("当前区块哈希值：%x\n", block.Hash)
	fmt.Printf("当前区块数据：%s\n", block.Data)
}
