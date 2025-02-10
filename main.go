package main

import (
	"fmt"

	"github.com/sodesu2077/matasano-crypto-go/internal/set1"
)

func main() {
	fmt.Println(string(set1.DetectAES("./static/data/set1.08.txt")))
}
