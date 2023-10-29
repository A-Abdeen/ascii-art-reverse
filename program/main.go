package main

import (
	"asciiart"
	"fmt"
	"os"
)

func main() {
	var rawInput string
	var outputFile string
	var standardfile []byte
	// Check if input is correct
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	} 
	if len(os.Args) < 2 {
		fmt.Println("Input string missing")
		return
	}
	// file := os.Args[1] // Read char file & string argument
	if len(os.Args) == 2 {
		if len(os.Args[1]) > 9 && os.Args[1][:10] == "--reverse=" {
			outputFile = os.Args[1][10:]
		} else {
			rawInput = os.Args[1]
		}
	}
	sourceFile, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i:=0;i<len(sourceFile);i++{
		if sourceFile[i] != 13{
			standardfile = append(standardfile, sourceFile[i])
		}
	}
	if rawInput != "" {
	// Main function: Splitting (split string based on newline position)
	// ∟--> Sub function: Formatting (change input to allow use of newline & qoutation marks)
	splitInput := asciiart.LineSplitter(rawInput, asciiart.InputFormatter)

	// Main function: Printing (printing the row of characters within input string)
	// ∟--> Sub function: Parsing (parsing the data of the 8 rows to print sequentially)
	asciiart.RowPrinter(splitInput, standardfile, asciiart.RowParser)} else {
		output, err := os.ReadFile(outputFile)
		if err != nil {
			fmt.Println(err)
			return
		}
	asciiart.FindWord(standardfile, output)
		
	}

}
