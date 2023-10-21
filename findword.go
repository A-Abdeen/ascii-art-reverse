package asciiart

import "fmt"

func FindWord(sourcefile []byte, output []byte) {
	var widthOfString int
	var widthOfletter int
	var count int
	var widthOfAllLetters []int
	var newoutput []byte
	var newLine int
	var nextLetter int
	var oneLetter []byte
	var finalLetter int
	var numberofLines int
	for i := 0; i < len(output); i++ {
		if output[i] == 10 {
			newLine++
		}
		if output[i] != 36 && output[i] != 13 {
			newoutput = append(newoutput, output[i])
		}
	}
	for h := 0; h < (newLine / 9); h++ {
		if h > 0 {
			fmt.Print("\\")
			fmt.Print("n")
		}
		for i := 0; i < len(newoutput); i++ {
			if newoutput[i] == 10 && widthOfString == 0 {
				widthOfString = i
			}
		}
		for i := 0; i < widthOfString; i++ {
			if newoutput[i] == 32 {
				for j := i; j < len(newoutput); {
					if newoutput[j] != 32 {
						count = 0
						break
					} else {
						count++
						j = j + widthOfString + 1
						if count == 8 {
							if widthOfletter == i {
								widthOfletter = widthOfletter + 6
								i = i + 5
							} else {
								widthOfletter = i + 1
							}
							widthOfAllLetters = append(widthOfAllLetters, widthOfletter)
							count = 0
							break
						}
					}
				}
			}
		}
		for i := 0; i < len(widthOfAllLetters); i++ {
			for k := 0; k <= 7; k++ {
				for j := 0; j < (widthOfAllLetters[i] - nextLetter); j++ {
					l := j + (k * (widthOfString + 1)) + nextLetter
					if l < len(newoutput) {
						oneLetter = append(oneLetter, newoutput[l])
					}
				}
				oneLetter = append(oneLetter, 10)
			}
			for m := 0; m < len(sourcefile); m++ {
				if sourcefile[m] == 10 {
					numberofLines++
				}
				for n := 0; n < len(oneLetter); n++ {
					if m+n > len(sourcefile)-1 {
						break
					}
					if oneLetter[n] != sourcefile[m+n] {
						break
					}
					if n > len(oneLetter)-2 {
						finalLetter = ((numberofLines) / 9) + 32
					}
				}
			}
			count = 0
			numberofLines = 0
			fmt.Print(string(finalLetter))
			oneLetter = []byte{}
			nextLetter = widthOfAllLetters[i]
			// fmt.Println(widthOfAllLetters[i])
			// fmt.Println(nextLetter)
		}

	}
	fmt.Println()
}
