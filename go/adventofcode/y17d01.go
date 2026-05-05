package adventofcode

import "strings"

func y17d01(input string, part int) (sum int) {
	nums := []int{}
	halfway := len(input) / 2
	for _, n := range strings.Split(input, "") {
		nums = append(nums, strToInt(n))
	}
	for i := 0; i < len(nums); i++ {
		check := (i + 1) % len(nums)
		if part == 2 {
			check = (i + halfway) % len(nums)
		}
		if nums[i] == nums[check] {
			sum += nums[i]
		}
	}
	return
}
