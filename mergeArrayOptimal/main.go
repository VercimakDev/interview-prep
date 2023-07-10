/* Merge two integer arrays in ascending order */
package main

import (
	"sort"
)

func main() {
	var nums1 = []int{1, 2, 3, 0, 0, 0}
	var nums2 = []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}
	sort.Ints(nums1)
}
