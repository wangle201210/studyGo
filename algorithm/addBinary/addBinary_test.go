package addBinary

import "testing"

func TestAddBinary(t *testing.T) {
	binary := addBinary("11", "111")
	t.Logf("data is: %+v", binary)
}
