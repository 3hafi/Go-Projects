package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// reading inputfile
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	words := strings.Fields(string(data))

	for _, word := range words {
		fmt.Println(word)
	}

	// identify patterns
	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" {
			hexValue := words[i-1]

			if _, err := strconv.ParseInt(hexValue, 16, 64); err != nil {
				continue
			}

			// converting hexvalue(string) to a decimal number
			decimalValue, err := strconv.ParseInt(hexValue, 16, 64)
			if err != nil {
				fmt.Println("Error processing HEX", err)
				return
			}
			words[i-1] = fmt.Sprintf("%d", decimalValue) //replacing previous word with dec value
			words = append(words[:i], words[i+1:]...)    // still unsure about, remove "(hex)" from the slice
			i--
		}
	}

	modifiedData := strings.Join(words, " ")
	fmt.Println("Modified data:", modifiedData)
	// writing outputfile
	err = os.WriteFile(os.Args[2], []byte(modifiedData), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File processing complete!")
}

/*
package main

import (
	"fmt"
	"io"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println(("incorrect argument, wanted 2 but got: "), len(os.Args))
	}

	if len(os.Args) ! = 2 {
		// err and use func isNotExist
	}
}
// error checking
// function to check correct os.Args commandline i.e filename

// function / err to check 2nd filename

//


/*modules:
1)Input/Output Handling: Read input from a file, write output to another file.

2)Text Processing: Perform modifications on the input text.
	- Hexadecimal to decimal conversion
	- Binary to decimal conversion
	- Case conversions (uppercase, lowercase, capitalized)
	- Punctuation formatting
	- Apostrophe handling
	- “A” to “An” conversion
	- Error Handling: Handle errors and edge cases

*/
