package wallet

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/sha256"
    "golang.org/x/crypto/ripemd160"
)

type Wallet struct {
    PrivateKey ecdsa.PrivateKey
    PublicKey  []byte
}

func NewWallet() *Wallet {
    private, public := newKeyPair()
    wallet := Wallet{private, public}

    return &wallet
}

func newKeyPair() (ecdsa.PrivateKey, []byte) {
    curve := elliptic.P256()
    private, _ := ecdsa.GenerateKey(curve, rand.Reader)
    public := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

    return *private, public
}

func PublicKeyHash(pubKey []byte) []byte {
    pubHash := sha256.Sum256(pubKey)

    hasher := ripemd160.New()
    _, _ = hasher.Write(pubHash[:])
    publicRIPEMD := hasher.Sum(nil)

    return publicRIPEMD
}