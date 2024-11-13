package sorting

import (
	"fmt"
	"strings"
)

// Function to print the binary tree level by level
func PrintHeap(heap []int) {
	if len(heap) == 0 {
		fmt.Println("Heap is empty")
		return
	}

	// Calculate the maximum height of the tree
	height := calculateHeight(len(heap))
	
	// Print each level of the tree
	for level := 0; level <= height; level++ {
		startIdx := (1 << level) - 1      // Starting index for this level
		endIdx := (1 << (level + 1)) - 2  // Ending index for this level
		if startIdx >= len(heap) {
			break
		}
		
		// Adjust the end index to avoid out-of-bounds access
		if endIdx >= len(heap) {
			endIdx = len(heap) - 1
		}
		
		// Print the nodes for this level
		levelValues := heap[startIdx:endIdx+1]
		fmt.Printf("Level %d: %v\n", level, levelValues)
	}
}

// Helper function to calculate the height of the tree
func calculateHeight(n int) int {
	height := 0
	for (1 << height) <= n {
		height++
	}
	return height - 1
}

// Function to print a visual representation of the binary tree
func PrintTree(heap []int) {
	// We calculate the height of the binary tree based on the number of elements in the heap
	height := calculateHeight(len(heap))
	width := (1 << (height + 1)) - 1  // Maximum width of the tree
	lines := make([]string, height+1)

	// Prepare each level of the tree
	for i := 0; i <= height; i++ {
		// Find the range of indices for the current level
		startIdx := (1 << i) - 1
		endIdx := (1 << (i + 1)) - 2
		if startIdx >= len(heap) {
			break
		}
		if endIdx >= len(heap) {
			endIdx = len(heap) - 1
		}
		
		// Build the line for the current level
		levelValues := heap[startIdx : endIdx+1]
		spacesBetweenNodes := (1 << (height - i + 1)) - 1
		line := ""
		for j, value := range levelValues {
			if j > 0 {
				// Add spaces between nodes
				line += strings.Repeat(" ", spacesBetweenNodes)
			}
			// Add the value of the node
			line += fmt.Sprintf("%d", value)
		}
		// Add padding to center-align the tree visually
		padLeft := (width - len(line)) / 2
		lines[i] = strings.Repeat(" ", padLeft) + line
	}

	// Print all levels with appropriate spacing between them
	for _, line := range lines {
		fmt.Println(line)
	}
}


