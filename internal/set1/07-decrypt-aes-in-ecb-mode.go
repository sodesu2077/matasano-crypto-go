package set1

import (
	"bufio"
	"crypto/aes"
	"fmt"
	"io"
	"os"

	"github.com/sodesu2077/matasano-crypto-go/utils"
)

func DecryptAES(fname string, key []byte) []byte {
	cipher, _ := aes.NewCipher(key)
	file, _ := os.Open(fname)
	defer file.Close()

	var buf []byte
	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading file", err)
			return nil
		}
		data, _ := utils.Base64ToBytes(line)
		buf = append(buf, data...)
	}

	for i := 0; i < len(buf); i += aes.BlockSize {
		cipher.Decrypt(buf[i:i+aes.BlockSize], buf[i:i+aes.BlockSize])
	}
	return buf
}
