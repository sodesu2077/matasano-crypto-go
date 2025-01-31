package set1

import (
	"fmt"
	"sort"
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

	data, key, err := FindSingleByteXOR(hexToBytes)
	fmt.Println("string before XOR: ", string(data), "\bXOR key: ", string(key))
	return data, key, err
}

func FindSingleByteXOR(data []byte) ([]byte, byte, error) {
	usage := make(map[byte]int)
	topFrequentBytes := make(map[byte]int)

	for _, b := range data {
		usage[b] += 1
		if usage[b] > 2 {
			topFrequentBytes[b] = usage[b]
		}
	}

	sortedTopFr := make([]byte, len(topFrequentBytes))
	for k := range topFrequentBytes {
		sortedTopFr = append(sortedTopFr, k)
	}
	sort.SliceStable(sortedTopFr, func(a, b int) bool {
		return topFrequentBytes[sortedTopFr[a]] > topFrequentBytes[sortedTopFr[b]]
	})

	for _, b := range sortedTopFr {
		for _, k := range FrequentChars {
			key := b ^ byte(k[0])
			result := make([]byte, len(data))
			for i, b := range data {
				result[i] = byte(key) ^ b
			}
			if len(result) != 0 && !strings.ContainsAny(string(result), "`@#$%^&*()_+=></~[|{}]��") {
				return result, key, nil
			}
		}
	}
	return nil, 0, nil
}
