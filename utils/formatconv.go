package utils

import "fmt"

const base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

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

func BytesToBase64(data []byte) (string, error) {
	/*
		Converting bytes to base64
		1. Iterate through the data slice
		2. Shift the buffer 8 bits to the left and add the new byte
		3. Increment the bit count by 8
		4. Process 6-bit chunks
			- Extract the last 6 bits from the buffer
			- Append the corresponding base64 character to the result
		5. Handle remaining bits (if any)
			- Pad the remaining bits with zeroes
			- Append the corresponding base64 character to the result
	*/

	var result string
	var buffer uint32
	var bitCount int

	for _, b := range data {
		buffer = (buffer << 8) | uint32(b)
		bitCount += 8

		for bitCount >= 6 {
			bitCount -= 6
			index := (buffer >> bitCount) & 0x3F
			result += string(base64Table[index])
		}
	}

	if bitCount > 0 {
		buffer <<= (6 - bitCount)
		result += string(base64Table[buffer&0x3F])
	}

	for len(result)%4 != 0 {
		result += "="
	}

	return result, nil
}