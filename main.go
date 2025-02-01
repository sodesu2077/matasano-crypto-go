package main

import (
	"fmt"

	"github.com/sodesu2077/matasano-crypto-go/challenges/set1"
)

func main() {
	// set1.HexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	// set1.FixedXOR("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	// set1.SingleByteXORCypher("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	// set1.DetectSingleCharXOR("./static/data/set1.04.txt")
	// set1.RepeatingKeyXOR("ICE", "Burning 'em, if you ain't quick and nimble I go crazy when I hear a cymbal")
	fmt.Println(string(set1.BreakRepeatingKeyXOR("./static/data/set1.06.txt")))
}
