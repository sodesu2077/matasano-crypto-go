package main

import (
	"fmt"

	"github.com/sodesu2077/matasano-crypto-go/challenges/set1"
)

func main() {
	// fmt.Println(string(set1.DecryptAES("./static/data/set1.07.txt", []byte("YELLOW SUBMARINE"))))
	fmt.Println(string(set1.DetectAES("./static/data/set1.08.txt")))
}
