//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重
//复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//
//
// 示例 2：
//
//
//输入：nums = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：nums = [0]
//输出：[]
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 3000
// -10⁵ <= nums[i] <= 10⁵
//
// Related Topics 数组 双指针 排序 👍 5051 👎 0

package leetcode

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	l := len(nums)
	for i := 0; i < l; i++ {
		//解答失败: 测试用例:[-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6]
		//测试结果:[[-4,-2,6],[-4,0,4],[-4,1,3],[-4,2,2],[-2,-2,4],[-2,0,2],[-2,-2,4],[-2,0,2]]
		//期望结果:[[-4,-2,6],[-4,0,4],[-4,1,3],[-4,2,2],[-2,-2,4],[-2,0,2]] stdout:
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					if len(ans) > 0 && ans[len(ans)-1][0] == nums[i] && ans[len(ans)-1][1] == nums[j] && ans[len(ans)-1][2] == nums[k] {
						continue
					}
					ans = append(ans, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
