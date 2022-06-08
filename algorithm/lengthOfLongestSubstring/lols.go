package lengthOfLongestSubstring

func lengthOfLongestSubstring(s string) int {
	max := 0
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		delete(m, s[i])
		for j := i + 1; j < len(s); j++ {
			// 存在就计算长度并且结束
			if m[s[j]] != 0 {
				max = maxInt(max, len(m))
				break
			}
			m[s[j]] = 1
			max = maxInt(max, len(m))
		}
	}
	max = maxInt(max, len(m))
	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func convert(s string, numRows int) string {
	arr := getArr(s, numRows)
	res := ""
	for _, v := range arr {
		res += v
	}
	return res
}

func getArr(s string, numRows int) (res []string) {
	for i := 0; i < len(s); {
		for j := 0; j < numRows; j++ {
			if i > len(s) {
				return
			}
			res[j] += string(s[i])
			i++
		}
	}
	return
}

func getArr1(s string, numRows int) (res []string) {
	res = make([]string, numRows)
	for i := 0; i < len(s); {
		for j := 0; j < numRows && i > len(s)-1; j++ {
			if i > len(s)-1 {
				return
			}
			res[j] += string(s[i])
			i++
		}
		for j := numRows - 2; j >= 1; j-- {
			if i > len(s)-1 {
				return
			}
			res[j] += string(s[i])
			i++
		}
	}
	return
}
