package utils

import "fmt"

func BytestoBase16(bytes []byte) (string, error) {
	hex := ""

	for _, v := range bytes {
		hex += fmt.Sprintf("%02x", v)
	}

	return hex, nil
}

func Base64ToBytes(text []byte) ([]byte, error) {
	base64Table := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	decodeTable := make(map[byte]int)
	for i, c := range base64Table {
		decodeTable[byte(c)] = i
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
