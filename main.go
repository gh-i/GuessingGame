package main

import (
	"fmt"
	"github.com/gh-i/GuessingGame/cryptoHelper"
	"github.com/gh-i/GuessingGame/player"
)

const NumberOfRounds = 1

func main() {
	guessList := []string{"King", "Queen"}
	alice := player.New(guessList)
	bob := player.New(guessList)

	for i := 0; i < NumberOfRounds; i++ {
		aliceSig := alice.SignedGuess()
		aliceGuess := alice.Guess()
		alicePubKey := alice.GetPublicKey()


		bobSig := bob.SignedGuess()
		bobGuess := bob.Guess()
		bobPubKey := bob.GetPublicKey()

		// verify alice guess
		fmt.Println("Alice", alicePubKey, aliceGuess, aliceSig)
		if cryptoHelper.VerifyMessage(alicePubKey, aliceGuess, aliceSig)  {
			fmt.Println("ALICE SUCCESS!")
		} else {
			fmt.Println("ALICE FAILED!")
		}

		// verify bobs guess
		fmt.Println("Bob", bobPubKey, bobGuess, bobSig)

		if cryptoHelper.VerifyMessage(bobPubKey, bobGuess, bobSig)  {
			fmt.Println("BOB SUCCESS!")
		} else {
			fmt.Println("BOB FAILED!")
		}
	}
}