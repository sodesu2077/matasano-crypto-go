package set2

import "bytes"

func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padData := bytes.Repeat([]byte{byte(padding)}, padding)

	data = append(data, padData...)
	return data
}
