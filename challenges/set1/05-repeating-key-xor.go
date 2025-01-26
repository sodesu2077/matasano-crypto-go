package set1

import (
	"fmt"

	"github.com/sodesu2077/matasano-crypto-go/utils"
)

func RepeatingKeyXOR(key string, text string) ([]byte, error) {
	counter := 0
	cypheredBytes := make([]byte, len(text))

	for i := 0; i < len(text); i++ {
		cypheredBytes[i] = byte(text[i]) ^ byte(key[counter])
		if counter == 2 {
			counter = 0
		} else {
			counter++
		}
	}

	hex, _ := utils.ConvertBytesIntoBase16(cypheredBytes)
	fmt.Println(hex)

	return cypheredBytes, nil
}
