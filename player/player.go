package player

import (
	"crypto/ecdsa"
	"crypto/rand"
	"github.com/gh-i/GuessingGame/cryptoHelper"
	"math/big"
)

type PlayerI interface {
	SignedGuess() *cryptoHelper.Signature
	Guess() []byte
}

type Player struct {
	guessList []string
	privateKey *ecdsa.PrivateKey
	publicKey *ecdsa.PublicKey
}

func New(guessList []string) *Player {
	privateK := cryptoHelper.NewKeyPair()
	pubK := privateK.PublicKey
	return &Player{guessList: guessList, privateKey: privateK, publicKey: &pubK}
}

func (p *Player) SignedGuess() *cryptoHelper.Signature {
	guess := p.guess()
	sig, err := cryptoHelper.SignMessage(p.privateKey, guess)
	if err != nil {
		panic(err)
	}
	return sig
}

func (p *Player) Guess() []byte {
	return p.guess()
}

func (p *Player) GetPublicKey() *ecdsa.PublicKey {
	return p.publicKey
}

func (p *Player) guess() []byte {
	n := len(p.guessList)
	index, err := rand.Int(rand.Reader, big.NewInt(int64(n)))

	if err != nil {
		panic(err)
	}

	return []byte(p.guessList[index.Int64()])
}