package sorting

func Quicksort(intList []int) {
	if len(intList) <= 1 {
		return
	}
	lo := 0
	hi := len(intList) - 1

	qsort(intList, lo, hi)

}

func qsort(intList []int, lo, hi int) {
	if lo >= hi {
		return
	}
	j := partition(intList, lo, hi)
	qsort(intList, lo, j-1)
	qsort(intList, j+1, hi)
}

func partition(intList []int, lo, hi int) int {
	i := lo
	j := hi + 1
	currVal := intList[lo]
	for {
		for {
			i = i + 1
			if currVal < intList[i] || i == hi {
				break
			}
		}
		for {
			j = j - 1
			if currVal > intList[j] || j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		Exch(intList, i, j)
	}
	Exch(intList, lo, j)
	return j
}


