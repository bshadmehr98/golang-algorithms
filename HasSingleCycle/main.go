package HasSingleCycle

import "fmt"

func GetCurrentIndex(i int, l int) int {
	if i < 0 {
		return l + i
	}
	return i
}

func HasSingleCycle(array []int) bool {
	currentIdx := 0
	currentIter := 0
	for range array {
		currentIdx = (currentIdx + array[currentIdx]) % len(array)
		currentIdx = GetCurrentIndex(currentIdx, len(array))
		if len(array) != 1 && array[currentIdx] == 0 {
			return false
		}
		if currentIdx == 0 && currentIter != len(array)-1 {
			return false
		}
		currentIter++
	}
	if currentIdx != 0 {
		return false
	}
	return true
}

func main() {
	a1 := []int{10, 11, -6, -23, -2, 3, 88, 909, -26}
	res := HasSingleCycle(a1)
	fmt.Println(res)
}
