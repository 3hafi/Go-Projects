package main

import (
	"bufio" // create buffer reader to read the file
	"fmt"   // to print output
	"math"
	"os" // to open the file
	"sort"
	"strconv" // to convert string to int
)

func readNumbersFromFile(filePath string) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file) // buffer to read the data into
	for scanner.Scan() {              // read the file line by line
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("invalid number in file: %v", err)
		}
		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return numbers, nil
}

func calculateAvergae(numbers []int) float64 {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}

func calculateMedian(numbers []int) float64 {
	sort.Ints(numbers)
	return float64(numbers[len(numbers)/2])
}

func calculateVariance(numbers []int) float64 {
	mean := calculateAvergae(numbers)
	sumSquaredDeviations := 0.0

	for _, num := range numbers {
		deviation := float64(num) - mean
		sumSquaredDeviations += deviation * deviation
	}
	return sumSquaredDeviations / float64(len(numbers))
}

func standardDeviation(numbers []int) float64 {
	return math.Sqrt(calculateVariance(numbers))
}

// Main function
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run your-program.go <file-path>")
		return
	}

	// Read numbers from file
	filePath := os.Args[1]
	numbers, err := readNumbersFromFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Calculate statistics
	average := calculateAvergae(numbers)
	median := calculateMedian(numbers)
	variance := calculateVariance(numbers)
	stdDev := standardDeviation(numbers)

	fmt.Println(numbers)
	fmt.Printf("Average: %.0f\n", average)
	fmt.Printf("Median: %.0f\n", median)
	fmt.Printf("Variance: %.0f\n", variance)
	fmt.Printf("Standard Deviation: %.0f\n", stdDev)
}
