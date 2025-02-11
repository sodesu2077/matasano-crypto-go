package set1

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/sodesu2077/matasano-crypto-go/utils"
)


func DecodeHexString(hex string) ([]byte, error) {
	strlen := len(hex)
	data := make([]byte, strlen/2)

	/*
		Converting hex string to bytes
		1. Iterate through the hex string in pairs, as each pair represents a byte (e.g., 1F = 00011111)
		2. Parse each hex pair into a byte using strconv.ParseUint
		3. Store the byte in the data slice
	*/
	for i := 0; i < strlen; i += 2 {
		pair := hex[i : i+2]
		b, err := strconv.ParseUint(pair, 16, 8)

		if err != nil {
			return []byte{}, err
		}

		data[i/2] = byte(b)

	}

	return data, nil
}



func HexToBase64(hex string) (string, error) {
	strlen := len(hex)
	if strlen%2 != 0 {
		return "", errors.New("invalid hex string provided")
	}
	data, err := DecodeHexString(hex)

	if err != nil {
		return "", err
	}

	result, _ := utils.BytesToBase64(data)

	fmt.Println(result)
	return result, nil
}
