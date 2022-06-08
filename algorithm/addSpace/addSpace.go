package addSpace

import "fmt"

func addSpaces(s string, spaces []int) string {
	b := []byte(s[:])
	for k, v := range spaces {
		index := v + k
		fmt.Println(string(b), "||||", string(b[index:]))
		temp := append(b[:index], ' ')
		temp = append(temp, b[index:]...)
		fmt.Println(string(b), "||||", string(b[index-1:]))
		b = temp
	}
	return string(b)
}
