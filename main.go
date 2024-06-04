package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	if len(arg) != 3 {
		fmt.Println("Usage: <go run .>", "<sample.txt>", "<result.txt>")
	}
}
