package ls

import (
	"fmt"
	"testing"
)

func TestLs(t *testing.T) {
	ls := NewLs(WithOpt("-R -A"), WithPath("/Users/med/mine/github/studyGo/cmd"))
	res := ls.Do()
	for _, r := range res {
		fmt.Printf("%+v\n", r)
		if len(r.Child) > 0 {
			for _, l := range r.Child {
				fmt.Printf("	%+v\n", l)
			}
		}
	}
}
