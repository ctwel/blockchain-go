package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// ProofOfWork 1. 定义一个结构
type ProofOfWork struct {
	// a. block
	block *Block
	// b. 目标值
	// 一个非常大的数
	target *big.Int
}

// NewProofOfWork 2. 创建POW的函数
func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block,
	}

	// 我们制定的难度值，现在是string类型，需要进行转换
	targetStr := "0000100000000000000000000000000000000000000000000000000000000000"
	// 引入的辅助变量，目的是将上面的难度值转为big.int
	tmpInt := big.Int{}
	// 将难度值赋值给big.int, 指定16进制的格式
	tmpInt.SetString(targetStr, 16)

	pow.target = &tmpInt
	return &pow
}

// Run 3. 提供计算pow的函数
func (pow *ProofOfWork) Run() ([]byte, uint64) {

	var nonce uint64
	var hash [32]byte
	block := pow.block

	for {
		// 1. 拼装数据（区块的数据，还有不断变化的随机数）
		tmp := [][]byte{
			Uint64ToByte(block.Version),
			block.PrevHash,
			block.MerkelRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			block.Data,
		}
		// 将二维的切片数组连接起来，返回一个一维的切片
		blockInfo := bytes.Join(tmp, []byte{})

		// 2. 做哈希运算
		hash = sha256.Sum256(blockInfo)

		// 3. 与pow的target进行比较
		// a. 找到了，推出返回
		// b. 没找到，继续找，随机数加1
		tmpInt := big.Int{}
		tmpInt.SetBytes(hash[:])
		if tmpInt.Cmp(pow.target) == -1 {
			// 找到了
			fmt.Printf("挖矿成功！ hash : %x, nonce : %d\n", hash, nonce)
			break
		} else {
			// 没找到
			nonce++
		}
	}

	//return []byte{}, 0
	return hash[:], nonce
}
