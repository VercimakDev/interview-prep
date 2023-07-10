package main
/* Two given integer array should be merged in ascending order */

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("hello world")
	var nums1 = []int {1,2,3,0,0,0}
	var nums2 = []int {2,5,6}
	merge(nums1, 3, nums2, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	num2Pointer := 0
	for i := m; i<m+n; i++{
		nums1[i] = nums2[num2Pointer]
		num2Pointer++
	}
	sort.Ints(nums1)
	fmt.Println(nums1)
}