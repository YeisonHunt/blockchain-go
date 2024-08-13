package blockchain

import (
    "bytes"
    "crypto/sha256"
    "encoding/gob"
    "time"
)

type Block struct {
    Timestamp    int64
    Transactions []*Transaction
    PrevHash     []byte
    Hash         []byte
    Nonce        int
}

func (b *Block) HashBlock() {
    var data bytes.Buffer
    encoder := gob.NewEncoder(&data)
    err := encoder.Encode(b)
    if err != nil {
        panic(err)
    }

    hash := sha256.Sum256(data.Bytes())
    b.Hash = hash[:]
}

func NewBlock(transactions []*Transaction, prevHash []byte) *Block {
    block := &Block{
        Timestamp:    time.Now().Unix(),
        Transactions: transactions,
        PrevHash:     prevHash,
        Hash:         []byte{},
        Nonce:        0,
    }

    block.HashBlock()
    return block
}

package blockchain

import (
    "bytes"
    "crypto/sha256"
    "math/big"
)

type ProofOfWork struct {
    Block  *Block
    Target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
    target := big.NewInt(1)
    target.Lsh(target, uint(256-16)) // Adjust difficulty here

    return &ProofOfWork{b, target}
}

func (pow *ProofOfWork) Run() (int, []byte) {
    var hashInt big.Int
    var hash [32]byte
    nonce := 0

    for {
        data := bytes.Join([][]byte{
            pow.Block.PrevHash,
            pow.Block.HashTransactions(),
            IntToHex(pow.Block.Timestamp),
            IntToHex(int64(nonce)),
        }, []byte{})

        hash = sha256.Sum256(data)
        hashInt.SetBytes(hash[:])

        if hashInt.Cmp(pow.Target) == -1 {
            break
        } else {
            nonce++
        }
    }

    return nonce, hash[:]
}