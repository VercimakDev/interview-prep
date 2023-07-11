package main


func main() {
	var nums = []int{0,1,2,2,3,0,4,2}
	val := 2
	removeElement(nums, val)
}

func removeElement(nums []int, val int) int {
	validPos := 0
	for i:= 0; i<len(nums); i++ {
		if nums[i] != val {
			nums[i], nums[validPos] = nums[validPos], nums[i]
			validPos++
		}
	}
	return validPos
}
