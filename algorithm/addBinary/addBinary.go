package addBinary

import "strconv"

func addBinary(a string, b string) string {
	la, lb, carry := len(a)-1, len(b)-1, 0
	ans := []int{}
	for la >= 0 || lb >= 0 {
		if la >= 0 {
			carry += int(a[la] - '0')
			la--
		}
		if lb >= 0 {
			carry += int(b[lb] - '0')
			lb--
		}
		ans = append(ans, carry%2)
		carry = carry / 2
	}
	if carry != 0 {
		ans = append(ans, carry)
	}
	return toString(ans)
}

func toString(b []int) (res string) {
	for _, i := range b {
		res = strconv.Itoa(i) + res
	}
	return res
}

// func revert (s string) string {
//     for i := 0; i < len(s)/2; i++ {
//         s[i],s[len(s)-1-i] = s[len(s)-1-i],s[i]
//     }
//     return s
// }
