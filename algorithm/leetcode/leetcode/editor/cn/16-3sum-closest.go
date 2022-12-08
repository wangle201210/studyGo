//给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
//
// 返回这三个数的和。
//
// 假定每组输入只存在恰好一个解。
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
//
//
// 示例 2：
//
//
//输入：nums = [0,0,0], target = 1
//输出：0
//
//
//
//
// 提示：
//
//
// 3 <= nums.length <= 1000
// -1000 <= nums[i] <= 1000
// -10⁴ <= target <= 10⁴
//
// Related Topics 数组 双指针 排序 👍 1215 👎 0

package leetcode

import (
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)

//解答失败: 测试用例:[-1,2,1,-4] 1 测试结果:1 期望结果:2 stdout:
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	little := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1
		sum := 0
		for j < k {
			sum = nums[i] + nums[j] + nums[k]
			if sum < target {
				j++
			}
			if sum > target {
				k--
			}
			if sum == target {
				return target
			}
		}
		little = min(little, sum, target)
	}
	return little
}

func min(i, j, target int) int {
	if abs(target-i) < abs(target-j) {
		return i
	}
	return j
}

func abs(i int) int {
	if i < 0 {
		i *= -1
	}
	return i
}

//leetcode submit region end(Prohibit modification and deletion)
