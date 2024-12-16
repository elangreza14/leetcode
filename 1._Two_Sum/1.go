package twosum

// 1. Two Sum

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, \
// and you may not use the same element twice.
// You can return the answer in any order.

// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

func twoSum(nums []int, target int) []int {
	// use hash map for this problem
	// res is target minus num
	// save the num for key and value for index
	// check if delta target - num, if exist just return in with the index
	// iterate every nums and store it in the hash map
	m := make(map[int]int)
	for i, num := range nums {
		res := target - num
		if v, ok := m[res]; ok {
			return []int{v, i}
		}
		m[num] = i
	}

	return nil
}
