package main

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
