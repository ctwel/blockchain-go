package main

import (
	"fmt"
)

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
