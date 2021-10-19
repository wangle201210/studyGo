package option

import (
	"strconv"
	"testing"
)

func TestPeople(t *testing.T) {
	tl := []struct {
		para People
		want string
	}{
		{
			para: People{name: "wanna", age: 25, height: 176, money: 9999999999},
			want: "Name:wanna, Age: 25, Height: 176, Money: 9999999999",
		},
		{
			para: People{name: "lisa", age: 24, height: 166, money: 99999999},
			want: "Name:lisa, Age: 24, Height: 166, Money: 99999999",
		},
		//....
	}
	for i, l := range tl {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			p := NewPeople(l.para.name, WithAge(l.para.age), WithHeight(l.para.height), WithMoney(l.para.money))
			s := p.String()
			if s == l.want {
				t.Logf("pass")
				return
			}
			t.Errorf("want (%s) got (%s)", l.want, s)
		})
	}
}
