package main

func istArmstrong(zahl int) bool {
	var result int
	var nums []int
	temp := zahl

	for temp > 0 { //add each digit into the slice nums
		digit := temp % 10
		nums = append(nums, digit)
		temp /= 10
	}
	for _, num := range nums {
		result += potenziere(num, len(nums))
	}
	return result == zahl
}
