package set1

import "fmt"

func BreakRepeatingKeyXOR(file string) {
	FindHammingDistance([]byte("wokka wokka!!!"), []byte("this is a test"))
}

/*
Finds hamming distance between 2 values
 1. Assume both values are of equal length
 2. XOR 2 bytes to get difference in bits
    E.g. If you wanna find distance between 0b01100011 and 0b10011010, you need to XOR these bytes first.
 3. Count 1s within the resulted XOR'd byte
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
	fmt.Println(count)
	return byte(count)
}
