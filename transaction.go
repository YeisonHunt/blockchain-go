package blockchain

import (
    "crypto/sha256"
)

type Transaction struct {
    ID     []byte
    Input  []byte
    Output []byte
}

func (t *Transaction) HashTransaction() {
    hash := sha256.Sum256(append(t.Input, t.Output...))
    t.ID = hash[:]
}

func NewTransaction(input, output []byte) *Transaction {
    transaction := &Transaction{Input: input, Output: output}
    transaction.HashTransaction()
    return transaction
}