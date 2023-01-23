package SmallestDifference

import (
	"fmt"
	"math"
	"sort"
)

func Abs(a int) int {
	fmt.Println(a)
	if a >= 0 {
		return a
	}
	return a * -1
}

func smallestDifference(array1, array2 []int) []int {
	a1, a2 := 0, 0
	b1, b2 := 0, 0
	res := make([]int, 2)
	mainDiff := math.MaxInt

	sort.Ints(array1)
	sort.Ints(array2)

	for a1 < len(array1) || a2 < len(array2) {
		dif1 := math.MaxInt
		dif2 := math.MaxInt
		if a1 < len(array1)-1 {
			dif1 = Abs(array1[a1+1] - array2[a2])
		}
		if a2 < len(array2)-1 {
			dif2 = Abs(array2[a2+1] - array1[a1])
		}

		if dif1 < dif2 {
			a1 += 1
			if dif1 <= mainDiff {
				b1 = a1
				b2 = a2
				mainDiff = dif1
			}
			if a1 > len(array1) {
				break
			}
		} else {
			a2 += 1
			if dif2 <= mainDiff {
				b2 = a2
				b1 = a1
				mainDiff = dif2
			}
			if a2 > len(array2) {
				break
			}
		}
	}
	res[0] = array1[b1]
	res[1] = array2[b2]
	return res
}

func main() {
	a1 := []int{-1, 5, 10, 20, 28, 3}
	a2 := []int{26, 134, 135, 15, 17}
	res := smallestDifference(a1, a2)
	fmt.Println(res)
}
