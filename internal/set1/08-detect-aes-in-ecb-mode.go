package set1

import (
	"bufio"
	"crypto/aes"
	"fmt"
	"io"
	"os"

	"github.com/sodesu2077/matasano-crypto-go/utils"
)

func DetectAES(fname string) []byte {
	f, err := os.Open(fname)
	utils.HandleErr(err)

	reader := bufio.NewReader(f)
	table := make(map[string]int)

	for {
		l, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading file", err)
			return nil
		}

		d, _ := DecodeHexString(string(l))

		for i := 0; i < len(d); i += aes.BlockSize {
			key := string(d[i : i+aes.BlockSize])
			table[key] += 1
			if table[key] >= 2 {
				return l
			}
		}
	}

	return nil
}
