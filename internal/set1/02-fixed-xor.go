package set1

import (
	"fmt"
	"strconv"

	"github.com/sodesu2077/matasano-crypto-go/utils"
)

func FixedXOR(buf1 string, buf2 string) (string, error) {
	buf1Data, err := DecodeHexString(buf1)

	utils.HandleErr(err)

	buf2Data, err := DecodeHexString(buf2)

	utils.HandleErr(err)

	result := ""

	for i := 0; i < len(buf1Data); i += 1 {
		xord := buf1Data[i] ^ buf2Data[i]
		value := strconv.FormatInt(int64(xord), 16)
		result += value
	}

	fmt.Println(result)

	return result, nil
}
