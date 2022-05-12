package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"time"
)

// Block 0. 定义结构
type Block struct {
	// 1. 版本号
	Version uint64
	// 2. 前区块hash
	PrevHash []byte
	// 3. Merkel 根
	MerkelRoot []byte
	// 4. 时间戳
	TimeStamp uint64
	// 5. 难度值
	Difficulty uint64
	// 6. 随机数
	Nonce uint64

	// a. 当前区块hash，正常比特币区块中没有当前区块的哈希，我们为了方便做了简化
	Hash []byte
	// b. 数据
	Data []byte
}

// Uint64ToByte 实现一个辅助函数，功能是将uint64 转为[]byte
func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}

// NewBlock 2. 创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Version:    00,
		PrevHash:   prevBlockHash,
		MerkelRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce:      0,
		Hash:       []byte{}, // 先填空，后面在计算 //TODO
		Data:       []byte(data),
	}
	block.SetHash()
	return &block
}

// SetHash 3. 生成Hash
func (block *Block) SetHash() {
	// 1. 拼装数据
	/*
		var blockInfo []byte
		blockInfo = append(blockInfo, Uint64ToByte(block.Version)...)
		blockInfo = append(blockInfo, block.PrevHash...)
		blockInfo = append(blockInfo, block.MerkelRoot...)
		blockInfo = append(blockInfo, Uint64ToByte(block.TimeStamp)...)
		blockInfo = append(blockInfo, Uint64ToByte(block.Difficulty)...)
		blockInfo = append(blockInfo, Uint64ToByte(block.Nonce)...)
		blockInfo = append(blockInfo, block.Data...)
	*/
	tmp := [][]byte{
		Uint64ToByte(block.Version),
		block.PrevHash,
		block.MerkelRoot,
		Uint64ToByte(block.TimeStamp),
		Uint64ToByte(block.Difficulty),
		Uint64ToByte(block.Nonce),
		block.Data,
	}

	// 将二维的切片数组连接起来，返回一个一维的切片
	blockInfo := bytes.Join(tmp, []byte{})

	// 2. sha256
	// func Sum256(data []byte) [Size]byte {
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
