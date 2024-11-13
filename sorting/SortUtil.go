package sorting

import (
	"fmt"
)

func IsSorted(intList []int) bool {
	for k := 0; k < len(intList)-1; k++ {
		if intList[k] > intList[k+1] {
			return false
		}
	}
	return true
}

func ViewList(intList []int) {
	for i := 0; i < len(intList); i++ {
		fmt.Println(intList[i])
	}
}

func Exch(intList []int, x, y int) {
	temp := intList[x]
	intList[x] = intList[y]
	intList[y] = temp
}
