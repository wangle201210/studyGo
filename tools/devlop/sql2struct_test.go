package devlop

import (
	"fmt"
	"testing"
)

func TestSql2struct(t *testing.T) {
	s := "-----"
	s2s := Sql2struct(s)
	fmt.Println(s2s.String())
}
