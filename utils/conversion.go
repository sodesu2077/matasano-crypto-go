package utils

import "fmt"

func ConvertBytesIntoBase16(bytes []byte) (string, error) {
	hex := ""

	for _, v := range bytes {
		hex += fmt.Sprintf("%02x", v)
	}

	return hex, nil
}
