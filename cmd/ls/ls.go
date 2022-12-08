package ls

import (
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"sort"
	"strings"
)

// 实现 linux 的 ls 命令

type Ls struct {
	path      string // 默认是当前目录
	recursion bool   // 递归列出子目录
	all       bool   // 显示全部文件，包括隐藏文件
	reverse   bool   // 倒序
}

type Option struct {
	f func(*Ls)
}

type lsRes struct {
	Filename string
	IsDir    bool
	Mode     os.FileMode
	Size     int64
	Child    []*lsRes
}

func defaultLs() *Ls {
	return &Ls{
		path: "./",
	}
}

func NewLs(opt ...*Option) *Ls {
	l := defaultLs()
	for _, o := range opt {
		o.f(l)
	}
	return l
}

func WithPath(path string) *Option {
	return &Option{
		f: func(ls *Ls) {
			ls.path = path
		},
	}
}

// WithOpt 兼容的命令
// -R 递归子目录
// -A、-a 展示全部 包含隐藏文件
// -r 将结果倒序
func WithOpt(s string) *Option {
	split := strings.Split(s, " ")
	return &Option{f: func(ls *Ls) {
		for _, sp := range split {
			switch sp {
			case "-R", "R":
				ls.recursion = true
			case "-A", "A", "-a", "a":
				ls.all = true
			case "-r", "r":
				ls.reverse = true
			}
		}
	}}
}

func (l *Ls) Do() (res []*lsRes) {
	dir, err := ioutil.ReadDir(l.path)
	if err != nil {
		panic(err)
	}
	for _, d := range dir {
		// 没有说显示全部 且是隐藏文件（夹）则不显示
		if !l.all && isHidden(d.Name()) {
			continue
		}
		d.Mode()
		d.Sys()
		r := &lsRes{Filename: d.Name(), IsDir: d.IsDir(), Mode: d.Mode(), Size: d.Size()}
		// 需要递归且是文件目录
		if l.recursion && d.IsDir() {
			child := copyLs(l)
			child.path = path.Join(l.path, d.Name())
			r.Child = child.Do()
		}
		res = append(res, r)
	}
	// 倒序
	if l.reverse {
		sort.Slice(res, func(i, j int) bool {
			return res[i].Filename > res[j].Filename
		})
	}
	return res
}

func isHidden(s string) bool {
	return len(s) == 0 || s[0] == '.'
}

// 复制一份指针里面的所有参数
func copyLs(m *Ls) *Ls {
	vt := reflect.TypeOf(m).Elem()
	newoby := reflect.New(vt)
	newoby.Elem().Set(reflect.ValueOf(m).Elem())
	return newoby.Interface().(*Ls)
}
