package set1

import (
	"bufio"
	"os"

	"github.com/sodesu2077/matasano-crypto-go/utils"
)

/*
Detects a phrase within a file and a single XOR byte against which it has been encrypted.
1. Read the file content
2. Reuse the code from set1.03
3. Return the original phrase(string) and the XOR character
*/
func DetectSingleCharXOR(fName string) ([]byte, byte, error) {
	file, err := os.Open(fName)
	utils.HandleErr(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		decyphered, key, err := SingleByteXORCypher(line)

		utils.HandleErr(err)

		if decyphered != nil {
			return decyphered, key, nil
		}
	}

	return nil, 0, nil
}
