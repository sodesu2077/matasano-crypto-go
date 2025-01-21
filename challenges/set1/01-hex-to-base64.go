package set1

import (
	"errors"
	"fmt"
	"strconv"
)

const base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func HexToBase64(hexStr string) (string, error) {
	strlen := len(hexStr)
	if strlen%2 != 0 {
		return "", errors.New("invalid hex string provided")
	}

	data := make([]byte, strlen/2)

	/*
	Converting hex string to bytes
	1. Iterate through the hex string in pairs, as each pair represents a byte (e.g., 1F = 00011111)
	2. Parse each hex pair into a byte using strconv.ParseUint
	3. Store the byte in the data slice
	*/
	for i := 0; i < strlen; i += 2 {
		pair := hexStr[i : i+2]
		b, err := strconv.ParseUint(pair, 16, 8)

		if err != nil {
			return "", err
		}

		data[i/2] = byte(b)

	}

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

	fmt.Println(result)
	return result, nil
}
