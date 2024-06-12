package calc

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	var list []int
	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err == nil {
			list = append(list, num)
		} else {
			n1, n2 := list[len(list)-2], list[len(list)-1]
			list = list[:len(list)-2]
			switch token {
			case "+":
				list = append(list, n1+n2)
			case "-":
				list = append(list, n1-n2)
			case "*":
				list = append(list, n1*n2)
			case "/":
				list = append(list, n1/n2)
			}
		}
	}
	return list[0]
}
