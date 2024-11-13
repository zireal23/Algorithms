package randomnumbergenerators

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/exp/rand"
)

func GenerateIntegers(N int) {
	rand.Seed(uint64(time.Now().UnixNano()))
	// Create or open the file for writing
	file, err := os.Create("./sorting/integers.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Generate 1000 integers and write each to a new line
	for i := 1; i <= N; i++ {
		randomNum := rand.Intn(100)
		fmt.Fprintf(file, "%d\n", randomNum) // Write each integer to a new line
	}

	fmt.Println("Successfully generated 1000 integers in integers.txt")
}
