package main

import (
	"fmt"       // Importing the fmt package for formatted input/output
	"os"        // Provides functions to work with operating system (files, args, etc.)
	"strconv"   // Provides conversion functions between strings and numbers
	"strings"   // Provides string manipulation functions (splitting, joining, etc.)
)

func main() {
	// Reading the input file (passed as the first command-line argument)
	data, err := os.ReadFile(os.Args[1]) // os.ReadFile reads the file content into a byte slice ([]byte)
	// os.Args[1] refers to the first command-line argument provided (the file path).
	// err captures any errors during file reading.
	if err != nil { // Check if an error occurred (err != nil means an error exists).
		fmt.Println(err) // Print the error message to the console.
		return           // Exit the program if there's an error.
	}

	// Splitting the file content into words
	words := strings.Fields(string(data)) // Converts byte slice to string and splits into a slice of words ([]string).
	// strings.Fields splits based on any whitespace (spaces, tabs, newlines).

	// Iterating over all words to print them
	for _, word := range words { // Loop through each word in the slice. '_' ignores the index.
		fmt.Println(word) // Print each word to the console.
	}

	// Identifying and processing patterns (e.g., "(hex)" and "(bin)")
	for i := 0; i < len(words); i++ { // Standard for loop iterating by index (i).
		if words[i] == "(hex)" { // Check if the current word is "(hex)".
			hexValue := words[i-1] // The preceding word is expected to be the hexadecimal value.

			// Validate the hexadecimal value (ensure it's valid before processing)
			if _, err := strconv.ParseInt(hexValue, 16, 64); err != nil {
				// strconv.ParseInt attempts to convert hexValue to a decimal integer (base 16).
				// The underscore `_` ignores the resulting value, as we only care about the error.
				continue // Skip this iteration if the hex value is invalid.
			}

			// Convert the hexadecimal string to a decimal number
			decimalValue, err := strconv.ParseInt(hexValue, 16, 64) // Convert valid hex to a decimal integer.
			if err != nil { // Handle any unexpected errors (rare if validation passed earlier).
				fmt.Println("Error processing HEX", err) // Print error and exit.
				return
			}

			// Replace the original hex value with the decimal equivalent
			words[i-1] = fmt.Sprintf("%d", decimalValue) // Format the decimalValue as a string and replace it in the slice.
			words = append(words[:i], words[i+1:]...)    // Remove the "(hex)" by combining parts of the slice.
			// words[:i] gives all elements before i.
			// words[i+1:] gives all elements after i.
			// The `...` expands the second slice into individual elements to append.
			i-- // Decrement the index to reprocess the slice after modification.
		}

		if words[i] == "(bin)" { // Check if the current word is "(bin)".
			binValue := words[i-1] // The preceding word is expected to be the binary value.

			// Validate the binary value
			if _, err := strconv.ParseInt(binValue, 2, 64); err != nil {
				// strconv.ParseInt attempts to convert binValue to a decimal integer (base 2).
				continue // Skip this iteration if the binary value is invalid.
			}

			// Convert the binary string to a decimal number
			decimalValue, err := strconv.ParseInt(binValue, 2, 64) // Convert valid binary to a decimal integer.
			if err != nil { // Handle any unexpected errors (rare if validation passed earlier).
				fmt.Println("Error processing BIN", err) // Print error and exit.
				return
			}

			// Replace the original binary value with the decimal equivalent
			words[i-1] = fmt.Sprintf("%d", decimalValue) // Format the decimalValue as a string and replace it in the slice.
			words = append(words[:i], words[i+1:]...)    // Remove the "(bin)" by combining parts of the slice.
			i-- // Decrement the index to reprocess the slice after modification.
		}
	}

	// Join the modified words back into a single string
	modifiedData := strings.Join(words, " ") // strings.Join combines all words with a single space between them.

	// Output the modified data
	fmt.Println("Modified data:", modifiedData) // Print the final modified data to the console.

	// Write the modified data to the output file (specified as the second command-line argument)
	err = os.WriteFile(os.Args[2], []byte(modifiedData), 0644) // os.WriteFile writes data to a file.
	// os.Args[2] is the output file path provided via the command line.
	// []byte(modifiedData) converts the string back to a byte slice for writing.
	// 0644 sets file permissions (-rw-r--r-- in Unix).
	if err != nil { // Check for errors during file writing.
		fmt.Println(err) // Print the error message if writing fails.
		return           // Exit the program.
	}

	fmt.Println("File processing complete!") // Indicate successful completion.
}
