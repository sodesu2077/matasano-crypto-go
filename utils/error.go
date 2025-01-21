package utils

import "fmt"

func HandleErr(err error) {
	if err != nil {
		defer panic(err)
		fmt.Println("process terminated:/")
	}
}
