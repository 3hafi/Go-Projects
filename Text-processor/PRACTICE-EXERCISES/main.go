package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
)

func main () {
	//reading the input-file
	data, err := os.ReadFile(os.Argss[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	words := strings.Fields(string(data))

	for _, word := range words {
		fmt.Println(word)		
	}

	//pattern recogonation
	for i := 0; i < len(words); i++ {
	//hex to sting

	if words[i] == "(hex)"
		//validation

	}





	//bin to sting#

	//modifiedData
}

