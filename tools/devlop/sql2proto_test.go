package devlop

import (
	"fmt"
	"testing"
)

func TestSql2proto(t *testing.T) {
	s := "------"
	s2p2 := Sql2proto(s, WithCommentPos("right"), WithGogo())
	fmt.Println(s2p2.String())
}
