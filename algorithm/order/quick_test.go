package order

import (
	"fmt"
	"testing"
)

func TestQuick(t *testing.T) {
	s := []int{1, 3, 5, 2, 6, 4, 8, 8, 2, 1, 9, 6, 3, 2, 3, 6}
	selec(s)
	fmt.Println(s)
}
