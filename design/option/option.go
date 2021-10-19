package option

import "fmt"

type People struct {
	name   string
	age    int64
	height int64
	money  int64
}

type Option struct {
	f func(p *People)
}

// 名字必须有，其他可以配置
func NewPeople(name string, opt ...*Option) *People {
	p := defaultOpt()
	p.name = name
	for _, option := range opt {
		option.f(p)
	}
	return p
}

func defaultOpt() *People {
	return &People{
		height: 175, // 不会还有人没有175吧
	}
}

func WithAge(age int64) *Option {
	return &Option{f: func(p *People) {
		p.age = age
	}}
}

func WithHeight(h int64) *Option {
	return &Option{f: func(p *People) {
		p.height = h
	}}
}

func WithMoney(m int64) *Option {
	return &Option{f: func(p *People) {
		p.money = m
	}}
}

func (p *People) String() string {
	return fmt.Sprintf("Name:%s, Age: %d, Height: %d, Money: %d", p.name, p.age, p.height, p.money)
}
