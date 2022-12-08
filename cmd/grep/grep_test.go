package grep

import "testing"

func TestGrep(t *testing.T) {
	grep := NewGrep()
	grep.regexp = "{$"
	grep.path = "/Users/med/mine/github/studyGo/cmd/grep"
	grep.Do()
}
