package main

import "math/big"

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

	return []byte{}, 0
}
