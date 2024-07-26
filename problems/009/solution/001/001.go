package main

func RemoveElement(nums []int, val int) int {
	j := 0

	for i := 0; i < len(nums)-1; i++ {

		if nums[i] != val {
			// se o valor for o mesmo, para não modificar
			if j != i {
				nums[j] = nums[i]
			}
			j++
		}
	}

	return j
}
