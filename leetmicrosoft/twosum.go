package leetmicrosoft

func twoSum(nums []int, target int) []int {

	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i, j)

				break
			}
		}
	}

	return result
}
