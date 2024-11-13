package disjointsets

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DisjointSettest() {
	fileName := os.Args[1]
	processUF(fileName)
}

func processUF(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fileReader := bufio.NewReader(file)

	// Read the first line to get the integer n
	firstLine, _, err := fileReader.ReadLine()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	n, err := strconv.Atoi(string(firstLine)) // Convert the first line to integer
	ds := NewDisjointSet(n)
	if err != nil {
		fmt.Println("Error converting first line to integer:", err)
		return
	}
	// Read the next n lines, each containing a pair of integers
	for i := 0; i < n; i++ {
		line, _, err := fileReader.ReadLine()
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}
		// Split the line into two integers
		parts := strings.Fields(string(line))
		if len(parts) != 2 {
			fmt.Println("Invalid pair format")
			continue
		}
		// Convert both parts to integers
		firstInt, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error converting first part to integer:", err)
			continue
		}
		secondInt, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error converting second part to integer:", err)
			continue
		}
		if !IsConnected(ds, firstInt, secondInt) {
			Union(ds, firstInt, secondInt)
			fmt.Printf("Pair: (%d, %d)\n", firstInt, secondInt)
		} else {
			fmt.Printf("Pair: (%d, %d) are already connected.\n", firstInt, secondInt)
		}
	}
}
