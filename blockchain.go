package blockchain

type Blockchain struct {
    Blocks []*Block
}

func NewBlockchain() *Blockchain {
    return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func (bc *Blockchain) AddBlock(transactions []*Transaction) {
    prevBlock := bc.Blocks[len(bc.Blocks)-1]
    newBlock := NewBlock(transactions, prevBlock.Hash)
    bc.Blocks = append(bc.Blocks, newBlock)
}

func NewGenesisBlock() *Block {
    return NewBlock([]*Transaction{}, []byte{})
}