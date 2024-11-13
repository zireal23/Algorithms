package sorting

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func TestMergeSorting() {
	intList := getSliceFromFile()
	MergeSortListTopDown(intList)
	fmt.Println(IsSorted(intList))
	ViewList(intList)
}

func TestQSorting() {
	intList := getSliceFromFile()
	Quicksort(intList)
	fmt.Println(IsSorted(intList))
	ViewList(intList)
}

func TestHeapSorting() {
	intList := getSliceFromFile()
	fmt.Println(intList)
	pq := NewPriorityQueueFromArray(intList)
	sortedArray := HeapSort(pq)
	fmt.Println("Sorted Array: ")
	fmt.Println(sortedArray)
}

func getSliceFromFile() []int {
	// Check if a filename is provided as a command-line argument
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file path as argument")
	}
	filePath := os.Args[1]
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	// Create a slice to hold the integers
	var numbers []int

	// Use bufio scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Convert each line to an integer
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Failed to convert line to integer: %s", err)
		}
		// Append the integer to the slice
		numbers = append(numbers, num)
	}

	// Check for errors in reading the file
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	return numbers
}
