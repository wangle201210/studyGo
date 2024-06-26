package life

import (
	"strconv"
	"testing"
)

func TestHouse(t *testing.T) {
	tl := []struct {
		data    []int64
		wantOld int64
		wantNew int64
	}{
		{
			data:    []int64{1000, 1500, 1500},
			wantOld: 4000 * 20,
			wantNew: (1000*3 + 1500*2 + 1500*1) * 9 / 10, // 整百才能这样算哦
		},
	}
	for i, s := range tl {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			fund := NewHosingFund(s.data)
			oldH := fund.HouseOld()
			if oldH != s.wantOld {
				t.Errorf("data is (%+v) wantOld (%d) got (%d)", s.data, s.wantOld, oldH)
			}
			newH := fund.HouseNew()
			if newH != s.wantNew {
				t.Errorf("data is (%+v) wantNew (%d) got (%d)", s.data, s.wantOld, newH)
			}
		})
	}
}

func TestHouse2(t *testing.T) {
	fund := NewHosingFund([]int64{})
	more, str := fund.MoreUse()
	println(more, str)
}
