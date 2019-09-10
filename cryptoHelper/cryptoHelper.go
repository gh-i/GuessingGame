package cryptoHelper

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
)

type Signature struct {
	R, S *big.Int
}

func (sig Signature) Signature() []byte {
	signature := sig.R.Bytes()
	return append(signature, sig.S.Bytes()...)
}

func NewKeyPair() *ecdsa.PrivateKey {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}
	return private
}

func SignMessage(privateK *ecdsa.PrivateKey, message []byte) (*Signature, error) {
	r, s, e := ecdsa.Sign(rand.Reader, privateK, message[:])
	return &Signature{r, s}, e
}

func VerifyMessage(pub *ecdsa.PublicKey, hash []byte, sig *Signature) bool {
	return ecdsa.Verify(pub, hash, sig.R, sig.S)
}