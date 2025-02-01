package set1

import (
	"bufio"
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

	bytes, err := utils.Base64ToBytes(string(text))
	utils.HandleErr(err)

	for KEYSIZE := 2; KEYSIZE <= 40; KEYSIZE++ {
		avg := DistanceAverage(bytes, KEYSIZE)
		if avg <= 1 {
			ptr := 0
			KEY := make([]byte, KEYSIZE)

			for ptr < KEYSIZE {
				block := make([]byte, len(bytes)/KEYSIZE)
				index := 0
				for j := ptr; j < len(bytes)-KEYSIZE-1; j += KEYSIZE {
					block[index] = bytes[j]
					index++
				}
				data, k, _ := FindSingleByteXOR(block)
				if data != nil {
					KEY[ptr] = k
				}
				ptr++
			}
			if KEY[0] != 0 {
				return KEY
			}
		}
	}
	return nil
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
