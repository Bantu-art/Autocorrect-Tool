package main

import (
	"fmt"
	"os"
	"io"
	"Autocorrect/Auto"
)

func main() {
	arg := os.Args
	if len(arg) != 3 {
		fmt.Println("Usage: <go run .>", "<sample.txt>", "<result.txt>")
		return
	}

	input := arg[1]

	// Open the input file
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the content of the input file
	fileContent, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Call the Auto.Modify function
	modified := Auto.Modify(string(fileContent))

	// Write to the output file
	output := arg[2]

	err = os.WriteFile(output, []byte(modified), 0o644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Print the modified content
	fmt.Println(modified)
}
