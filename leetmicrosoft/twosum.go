package leetmicrosoft

func twoSum(nums []int, target int) []int {
	values := make(map[int][]int)

	for i := 0; i < len(nums); i++ {
		indexes, ok := values[nums[i]]

		if ok {
			values[nums[i]] = append(indexes, i)
		} else {
			values[nums[i]] = []int{i}
		}

	}

	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		rest := target - nums[i]

		indexes, ok := values[rest]

		if ok {
			for _, j := range indexes {
				if j != i {
					result = append(result, i, j)
					break
				}
			}
		}

		if len(result) != 0 {
			break
		}
	}

	return result
}
