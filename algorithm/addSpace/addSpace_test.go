package addSpace

import "testing"

func TestAddSpace(t *testing.T) {
	spaces := addSpaces("LeetcodeHelpsMeLearn", []int{8, 13, 15})
	t.Log(spaces)
}
