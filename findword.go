package asciiart

import "fmt"

func FindWord(sourcefile []byte, output []byte) {
	var newoutput []byte
	var newLine int
	for i := 0; i < len(output); i++ {   
		if output[i] == 10 {     // count number of lines to find out if you need \n
			newLine++
		}
		if output[i] != 36 && output[i] != 13 { // remove 36($) and remove 13 new line in windows
			newoutput = append(newoutput, output[i])
		}
	}
	for h := 0; h < (newLine / 9); h++ { // if there are more than 17 new lines i means there is \n
	var nextLetter int
	var oneLetter []byte
	var finalLetter int
	var numberofLines int
	var widthOfString int
	var widthOfletter int
	var count int
	var widthOfAllLetters []int
		if h > 0 { // if it is the first time over the loop no need for \n 
			fmt.Print("\\")
			fmt.Print("n")
		}
		for i := 0; i < len(newoutput); i++ {
			if newoutput[i] == 10 && widthOfString == 0 {
				widthOfString = i
			}
		}
		for i := 0; i < widthOfString; i++ {
			if newoutput[i] == 32 { // if a space is found check the next 8 lines if there are all spaces means it is one character
				for j := i; j < len(newoutput); {
					if newoutput[j] != 32 {
						count = 0
						break
					} else {
						count++
						j = j + widthOfString + 1 // loop over going to the spot below each time 
						if count == 8 {
							if widthOfletter == i { // to check for a space character used
								widthOfletter = widthOfletter + 6 
								i = i + 5
							} else { 
								widthOfletter = i + 1 // width of letter determines when the character ends
							}
							widthOfAllLetters = append(widthOfAllLetters, widthOfletter) // append the end of each letter in a slice of int
							count = 0
							break
						}
					}
				}
			}
		}
		for i := 0; i < len(widthOfAllLetters); i++ { // go over the slice with the ends of each character
			for k := 0; k <= 7; k++ { // loop to append the data of the entire character as an array of byte
				for j := 0; j < (widthOfAllLetters[i] - nextLetter); j++ {
					l := j + (k * (widthOfString + 1)) + nextLetter // to move to the spot below it exactly
					if l < len(newoutput) {
						oneLetter = append(oneLetter, newoutput[l])
					}
				}
				oneLetter = append(oneLetter, 10) // after appending the data of each line of the character append \n
			}
			for m := 0; m < len(sourcefile); m++ { // after getting the data as a slice we need to match it with the source file
				if sourcefile[m] == 10 { // to find which line you reached
					numberofLines++
				}
				for n := 0; n < len(oneLetter); n++ { // loop over the output file you want to read
					if m+n > len(sourcefile)-1 {
						break
					}
					if oneLetter[n] != sourcefile[m+n]{ // if they dont match move to the next
						break
					}
					if n > len(oneLetter)-3 { // if they matched all the way through the file you want to read get the ascii value of the character 
						finalLetter = ((numberofLines) / 9) + 32 
					}
				}
			}
			count = 0
			numberofLines = 0
			fmt.Print(string(finalLetter)) // print the letter found and go over to the next letter
			oneLetter = []byte{}
			nextLetter = widthOfAllLetters[i]
		}
	}
	fmt.Println()
}
