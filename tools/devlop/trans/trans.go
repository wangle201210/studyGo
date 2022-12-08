package trans

import "strconv"

type Base struct {
	Str  string
	From int64
	To   int64
	Ch   bool
}

func (b *Base) Trans() {
	parseInt, err := strconv.ParseInt(b.Str, 16, 64)
	if err != nil {
		panic(err)
	}
	println(parseInt)
}

func t8(s string) {
}
