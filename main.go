package main

import (
	"fmt"

	"github.com/sodesu2077/matasano-crypto-go/internal/set2"
)

func main() {
	fmt.Printf("%0x\n", set2.PKCS7Padding([]byte("YELLOW SUBMARINE"), 20))
}
