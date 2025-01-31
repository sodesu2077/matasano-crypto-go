package set1

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/sodesu2077/matasano-crypto-go/utils"
)

func BreakRepeatingKeyXOR(fname string) []byte {
	f, err := os.Open(fname)
	utils.HandleErr(err)

	defer f.Close()
	reader := bufio.NewReader(f)
	var text []byte

	for {
		line, _, err := reader.ReadLine()
		text = append(text, line...)
		if err == io.EOF {
			break
		}
	}

	bytes, err := DecodeBase64IntoBinary(string(text))
	utils.HandleErr(err)

	for KEYSIZE := 2; KEYSIZE <= 40; KEYSIZE++ {
		avg := DistanceAverage(bytes, KEYSIZE)
		if avg <= 1 {
			keySizePointer := 0
			KEY := make([]byte, KEYSIZE)

			for keySizePointer < KEYSIZE {
				block := make([]byte, len(bytes)/KEYSIZE)
				index := 0
				for j := keySizePointer; j < len(bytes)-KEYSIZE-1; j += KEYSIZE {
					block[index] = bytes[j]
					index++
				}
				data, key, _ := FindSingleByteXOR(block)
				if data != nil {
					KEY[keySizePointer] = key
				}
				keySizePointer++
			}
			if KEY[0] != 0 {
				fmt.Println(string(KEY))
				return KEY
			}
		}
	}
	return nil
}

func DistanceAverage(bytes []byte, KEYSIZE int) int {
	distances := make([]int, 5)
	for d := 1; d < len(distances); d++ {
		a := (d - 1) * KEYSIZE
		b := d * KEYSIZE
		c := (d + 1) * KEYSIZE
		distances[d] = int(FindHammingDistance(bytes[a:b], bytes[b:c])) / KEYSIZE
	}
	sum := 0
	for _, v := range distances {
		sum += v
	}
	return sum / len(distances)
}

func DecodeBase64IntoBinary(text string) ([]byte, error) {
	base64Table := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	decodeTable := make(map[rune]int)
	for i, c := range base64Table {
		decodeTable[c] = i
	}

	bytes := make([]byte, len(text)*6/8)
	var buffer uint32
	var bitCount int
	var index int

	for _, v := range text {
		buffer = (buffer << 6) | uint32(decodeTable[v])
		bitCount += 6

		for bitCount >= 8 {
			bytes[index] = byte(buffer >> (bitCount - 8))
			index++
			bitCount -= 8
		}
	}
	return bytes, nil
}

/*
Finds hamming distance between 2 values
 1. Assume both values are of equal length
 2. XOR 2 bytes to get difference in bits
    E.g. If you wanna find distance between 0b01100011 and 0b10011010, you need to XOR these bytes first.
 3. Count 1s within the resulted XOR'd byte using bitwise operations: left shift and XOR
*/
func FindHammingDistance(val1, val2 []byte) byte {
	count := 0
	for i, v := range val1 {
		xoredResult := v ^ val2[i]

		// Count 1s within the xored result byte
		distanceXorKey := 1

		for {
			if (xoredResult ^ byte(distanceXorKey)) < xoredResult {
				count++
			}
			if distanceXorKey == 128 {
				break
			}
			distanceXorKey <<= 1
		}
	}
	return byte(count)
}
