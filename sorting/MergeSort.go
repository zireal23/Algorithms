package sorting

import (
	"fmt"
	"strconv"
)

func mergeLists(intList []int, lo, mid, hi int) {
	numElements := len(intList)
	auxList := make([]int, numElements)

	for k := lo; k <= hi; k++ {
		auxList[k] = intList[k]
	}

	i := lo
	j := mid + 1

	//comparing both the side: i----mid and mid+1-----j
	//if left side element lesser or equal to right side element, then we put the left side element
	// else we put the right side element in the current position
	//if left side of array exhausts, populate the remaining right side of the array
	//if right side exhausts, then populate the remaining left side
	for k := lo; k <= hi; k++ {
		if i > mid {
			intList[k] = auxList[j]
			j++
		} else if j > hi {
			intList[k] = auxList[i]
			i++
		} else if auxList[i] <= auxList[j] {
			intList[k] = auxList[i]
			i++
		} else {
			intList[k] = auxList[j]
			j++
		}
	}
}

func MergeSortListBottomUp(intList []int) {
	lo := 0
	hi := len(intList) - 1
	fmt.Println("The length of the array is: " + strconv.Itoa(hi))
	msort(intList, lo, hi)
}

func MergeSortListTopDown(intList []int) {
	if len(intList) <= 1 {
		return
	}
	N := len(intList)
	for sz := 1; sz < N; sz += sz {
		for lo := 0; lo <= N-sz; lo += sz + sz {
			mergeLists(intList, lo, lo+sz-1, min(N-1, lo+sz+sz-1))
		}
	}

}

func msort(intList []int, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := lo + (hi-lo)/2
	msort(intList, lo, mid)
	msort(intList, mid+1, hi)
	mergeLists(intList, lo, mid, hi)
}
