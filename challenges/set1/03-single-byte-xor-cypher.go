package set1

import (
	"fmt"
	"strings"

	"github.com/sodesu2077/matasano-crypto-go/utils"
)

var FrequentChars []string = []string{"e", "t", "a", "o", "i", "n", "s", "h", "r", "d", "l", "u"}

/*
	Decrypts a message (hex string); Finds a single character (key) against which a hex string has been XOR'd
*/

func SingleByteXORCypher(hex string) ([]byte, byte, error) {
	hexToBytes, err := DecodeHexString(hex)
	utils.HandleErr(err)

	usage := make(map[byte]byte)
	topFrequentBytes := make(map[byte]byte)

	for _, b := range hexToBytes {
		usage[b] += 1
		if usage[b] > 2 {
			topFrequentBytes[b] = usage[b]
		}
	}

	for b, _ := range topFrequentBytes {

		for _, k := range FrequentChars {
			key := b ^ byte(k[0])
			var result []byte
			for _, b := range hexToBytes {
				result = append(result, byte(key)^b)
			}
			if !strings.ContainsAny(string(result), "`@#$%^&*()_+=></~[|{}]�") {
				fmt.Println("string before XOR: ", string(result), "\bXOR key: ", string(key))
				return result, key, nil
			}
		}
	}
	return nil, 0, nil
}
