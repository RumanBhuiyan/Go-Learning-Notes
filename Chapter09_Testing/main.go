package main

func Average(nums []int32) float32 {
	summation := int32(0)
	for _, value := range nums {
		summation += value
	}
	return float32(summation) / float32(len(nums))
}
