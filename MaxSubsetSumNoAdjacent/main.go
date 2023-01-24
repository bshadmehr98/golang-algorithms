package MaxSubsetSumNoAdjacent

import "fmt"

func MaxSubsetSumNoAdjacent(array []int) int {
	//a := make([]int, len(array))
	for i, v := range array {
		if i < 2 {
			array[i] = v
		} else if i == 2 {
			array[i] = v + array[0]
		} else {
			if array[i-2] > array[i-3] {
				array[i] = v + array[i-2]
			} else {
				array[i] = v + array[i-3]
			}
		}
	}
	if len(array) >= 2 {
		if array[len(array)-1] > array[len(array)-2] {
			return array[len(array)-1]
		} else {
			return array[len(array)-2]
		}
	} else if len(array) == 1 {
		return array[0]
	}
	return 0
}

func main() {
	//a1 := []int{75, 105, 120, 75, 90, 135}
	a1 := []int{75, 105, 120, 75, 90, 135}
	res := MaxSubsetSumNoAdjacent(a1)
	fmt.Println(res)
}
