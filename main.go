package main

func main() {
}

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		b := target - v
		n := m[b]

		if n != i && nums[n]+v == target {
			return []int{n, i}
		}
		m[v] = i
	}

	return []int{}
}
