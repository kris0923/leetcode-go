/*
 * 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
 */
package twoSum

func twoSum(nums []int, target int) []int {
	var result = []int{-1, -1}
	var length = len(nums)
	var numMap map[int]int
	numMap = make(map[int]int)
	for i := 0; i < length; i++ {
		key := target - nums[i]
		if _, ok := numMap[key]; ok {
			result[0] = numMap[key]
			result[1] = i
			return result
		} else {
			numMap[nums[i]] = i
		}
	}
	return result
}
