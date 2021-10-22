/*
10阶楼梯
每次1-3阶
3、6不允许走
一共有多少种走法
*/
package step

import "fmt"

type step [][]int

func Answer() {
	s := f(10)
	for i, j := range s {
		fmt.Printf("%d:%+v\n", i, j)
	}
	s2 := f1(10)
	fmt.Printf("步数%d\n", s2)
}

// 枚举出来
func f(x int) (s step) {
	if x == 1 {
		s = [][]int{{1}}
		return
	}
	if x == 2 {
		s = [][]int{{1, 1}, {2}}
		return
	}
	if x == 3 {
		//s = [][]int{{1, 1, 1}, {1, 2}, {2, 1}, {3}}
		// 不能上第三阶
		s = [][]int{}
		return
	}
	if x == 6 {
		// 不能上第六阶
		s = [][]int{}
		return
	}
	for _, i := range f(x - 1) {
		s = append(s, append(i, 1))
	}
	for _, i := range f(x - 2) {
		s = append(s, append(i, 2))
	}
	for _, i := range f(x - 3) {
		s = append(s, append(i, 3))
	}
	return
}

// 仅计算总共方式
func f1(x int) int {
	if x == 1 {
		return 1
	}
	if x == 2 {
		return 2
	}
	if x == 3 {
		return 0
	}
	if x == 6 {
		return 0
	}
	return f1(x-1) + f1(x-2) + f1(x-3)
}
