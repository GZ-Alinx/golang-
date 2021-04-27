package main

import (
	"fmt"
	"sort"
)

func main() {
	// 排序 针对int数据类型进行排序，使用sort模块进行排序

	// 整型
	nums := []int{55, 2, 5, 7, 4, 21, 1, 6, 90, 99}
	sort.Ints(nums)
	fmt.Println(nums)

	// 浮点型
	scores := []float64{5, 1.3, 8, 31, 22, 77, 9, 33.2, 21.1, 7, 9, 132, 12, 2, 36, 11, 2}
	sort.Float64s(scores)
	fmt.Println(scores)

	// 切片排序
	slice := [][2]int{{1, 2}, {2, 3}, {6, 9}, {4, 1}, {5, 8}}
	less := func(i, j int) bool {
		return slice[i][1] < slice[j][1]
	}
	sort.Slice(slice, less)
	fmt.Println(slice)
}
