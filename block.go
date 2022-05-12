package main

import "crypto/sha256"

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
