package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) == 3 {
		// Read Data from the file... okay if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
		Data := ReadFileToString(os.Args[1])

		// Identify positions of brackets
		startBrackets := indexOfStartBrackets(Data)
		endBrackets := indexOfEndBrackets(Data)
		totalBrackets := len(startBrackets)

		// Iterate through all bracketed instructions
		for i := 0; i < totalBrackets; i++ {
			// Extract Data inside brackets
			insideBrackets := Data[startBrackets[i]+1 : endBrackets[i]]
			// Get the transformation and number
			transformation, _ := returnSubStrAndNum(insideBrackets)
			// Apply corresponding transformations
			switch transformation {
			case "hex":
			Data = hexToDecimal(Data)
			case "bin":
			Data = binaryToDecimal(Data)
			case "up":
				Data = up(Data)
			case "low":
				Data = low(Data)
			case "cap":
				Data = cap(Data)
			}
			// Recheck brackets for next iteration
			startBrackets = indexOfStartBrackets(Data)
			endBrackets = indexOfEndBrackets(Data)
		}

		// Clean up brackets and apply additional transformations
		Data = removeBracketsData(Data)
		Data = transformAToAn(Data)
		Data = fPunctuation(Data)
		Data = fPunctuation2(Data)

		// Output the result and save it to a new file
		fmt.Println("\n", Data)
		StringToWriteFile(os.Args[2], Data)
	}
}

func ReadFileToString(s string) string {
	r, err := os.ReadFile(os.Args[1]) // just pass the file name
	if err != nil {
		log.Fatal(err)
	}
	return string(r) // convert Data to a 'string'
}

func StringToWriteFile(filename, myString string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	_, err2 := f.WriteString(myString)
	if err2 != nil {
		fmt.Println(err2)
	}
}

func hexToDecimal(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" && i > 0 {
			hexValue := words[i-1]
			decimalNum, err := strconv.ParseInt(hexValue, 16, 64)
			if err == nil {
				words[i-1] = fmt.Sprintf("%d", decimalNum)
			}
		}
	}

	return strings.Join(words, " ")
}

func binaryToDecimal(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		if words[i] == "(bin)" && i > 0 {
			binValue := words[i-1]
			decimalNum, err := strconv.ParseInt(binValue, 2, 64)
			if err == nil {
				words[i-1] = fmt.Sprintf("%d", decimalNum)
			}
		}
	}

	return strings.Join(words, " ")
}

func up(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], "(up)") {
			startIndex := i - 1
			for j := startIndex; j < i; j++ {
				words[j] = strings.ToUpper(words[j])
			}
		}
		if strings.HasPrefix(words[i], "(up,") {
			parts := strings.Split(words[i+1], ")")
			if len(parts) >= 2 {
				number, err := strconv.Atoi(strings.TrimSuffix(parts[0], " "))
				if err == nil {
					startIndex := i - number
					for j := startIndex; j < i; j++ {
						words[j] = strings.ToUpper(words[j])
					}
				}
			}
		}
	}
	return strings.Join(words, " ")
}

func low(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], "(low)") {
			startIndex := i - 1
			for j := startIndex; j < i; j++ {
				words[j] = strings.ToLower(words[j])
			}
		}
		if strings.HasPrefix(words[i], "(low,") {
			parts := strings.Split(words[i+1], ")")
			if len(parts) >= 2 {
				number, err := strconv.Atoi(strings.TrimSuffix(parts[0], " "))
				if err == nil {
					startIndex := i - number
					for j := startIndex; j < i; j++ {
						words[j] = strings.ToLower(words[j])
					}
				}
			}
		}
	}
	return strings.Join(words, " ")
}

func cap(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], "(cap)") {
			startIndex := i - 1
			// Loop through the words and capitalize them
			for j := startIndex; j < i; j++ {
				words[j] = strings.Title(words[j])
			}
		}
		if strings.HasPrefix(words[i], "(cap,") {
			parts := strings.Split(words[i+1], ")")
			if len(parts) >= 2 {
				// strings.TrimSuffix removes the space at the end of the string
				number, err := strconv.Atoi(strings.TrimSuffix(parts[0], " "))
				if err == nil {
					startIndex := i - number
					for j := startIndex; j < i; j++ {
						words[j] = strings.Title(words[j])
					}
				}
			}
		}
	}
	return strings.Join(words, " ")
}

func removeBracketsData(input string) string {
	var result string
	inBracket := false

	for _, char := range input {
		if char == '(' {
			inBracket = true
		} else if char == ')' {
			inBracket = false
		} else if !inBracket {
			// append outside characters
			result += string(char)
		}
	}

	return result
}

func fPunctuation(input string) string {
	var output []rune
	var prevRune rune

	for index, r := range input {

		if r == '.' || r == ',' || r == '!' || r == '?' || r == ':' || r == ';' {
			if prevRune == ' ' {
				// Remove the space before punctuation
				output = output[:len(output)-1]
			}
			output = append(output, r)
			if index != len(input)-1 {
				if !unicode.IsPunct(rune(input[index+1])) && rune(input[index+1]) != ' ' {
					output = append(output, ' ') // Add a space after punctuation
				}
			}

		} else {
			output = append(output, r)
		}

		prevRune = r
	}

	return string(output)
}

func fPunctuation2(input string) string {
	var output []rune
	var prevRune rune

	for _, r := range input {

		if r == '\'' || r == ' ' {
			if prevRune == ' ' && r == '\'' {
				// Remove the space before punctuation
				output = output[:len(output)-1]
				output = append(output, r)
			} else if prevRune == '\'' && r == ' ' {
			} else {
				output = append(output, r)
			}
		} else {
			output = append(output, r)
		}
		if r == ':' && (len(output) > 1 && output[len(output)-1] != ' ') {
			output = append(output, ' ')
		}
		prevRune = r
	}
	return string(output)
}

func returnSubStrAndNum(newStr string) (string, int) {
	if hasComma(newStr) {
		tempStr := strings.Split(newStr, " ")
		split, err := strconv.Atoi(tempStr[1])
		if err != nil {
			fmt.Println(err)
		}
		returnString := strings.Split(tempStr[0], ",")
		return returnString[0], split
	} else {
		return newStr, 1
	}
}

func indexOfStartBrackets(s string) []int {
	var ind []int
	for index, v := range s {
		if v == '(' {
			ind = append(ind, index)
		}
	}
	return ind
}

func indexOfEndBrackets(s string) []int {
	var ind []int
	for index, v := range s {
		if v == ')' {
			ind = append(ind, index)
		}
	}
	return ind
}

func hasComma(s string) bool {
	for _, v := range s {
		if v == ',' {
			return true
		}
	}
	return false
}

func whatChangeAtoAn(word string) bool {
	if len(word) == 0 {
		return false
	}
	firstChar := rune(word[0])
	vowels := "aeiouAEIOU"
	return strings.ContainsRune(vowels, firstChar) || firstChar == 'h' || firstChar == 'H'
}

func transformAToAn(input string) string {
	words := strings.Fields(input) // Split the input into words
	result := make([]string, 0, len(words))

	inQuotes := false // Flag to track if we're inside apostrophes

	for i := 0; i < len(words); i++ {
		currentWord := words[i]
		// Check for words inside apostrophes (quoted parts)
		if strings.HasPrefix(currentWord, "'") && !strings.HasSuffix(currentWord, "'") {
			inQuotes = true
		}
		if inQuotes {
			if strings.HasSuffix(currentWord, "'") {
				inQuotes = false
			}
			result = append(result, currentWord)
			continue
		}
		// Check for standalone "a" or "A" and next word starting with a vowel or 'h'
		if (currentWord == "a" || currentWord == "A") && i+1 < len(words) {
			nextWord := words[i+1]
			if whatChangeAtoAn(nextWord) {
				if currentWord == "a" {
					currentWord = "an"
				} else {
					currentWord = "An"
				}
			}
		}
		result = append(result, currentWord)
	}
	return strings.Join(result, " ")
}
