package grep

import (
	"bufio"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"regexp"
	"strings"
)

type Grep struct {
	path      string // 默认是当前目录
	recursion bool   // 递归列出子目录
	all       bool   // 显示全部文件，包括隐藏文件
	regexp    string // 匹配规则
	//reverse   bool   // 倒序
}

type grepRes struct {
	Filename string // 文件名字
	Filepath string // 文件路径
	Number   int64  // 内容在第几行
	Content  string // 那行具体的内容
}

type Option struct {
	f func(*Grep)
}

func defaultGrep() *Grep {
	return &Grep{
		path: "./",
	}
}

func NewGrep(opt ...*Option) *Grep {
	l := defaultGrep()
	for _, o := range opt {
		o.f(l)
	}
	return l
}

func WithOpt(s string) *Option {
	split := strings.Split(s, " ")
	return &Option{f: func(g *Grep) {
		for _, sp := range split {
			switch sp {
			case "-R", "R":
				g.recursion = true
			case "-A", "A", "-a", "a":
				g.all = true
			}
		}
	}}
}

func (g *Grep) Do() (res []*grepRes) {
	return g.doFile()
}

func (g *Grep) doFile() (res []*grepRes) {
	fsrc, err := os.Open(g.path)
	if err != nil {
		panic(err)
	}
	defer fsrc.Close()
	if fsrc.Mode().IsDir() {
		// 不递归就不往后看了
		if !g.recursion {
			return
		}
		dir, err := ioutil.ReadDir(g.path)
		if err != nil {
			panic(err)
		}
		for _, d := range dir {
			n := copyGrep(g)
			n.path = path.Join(g.path, d.Name())
			res = append(res, n.doFile()...)
		}
		return
	}

	fileScanner := bufio.NewScanner(fsrc)
	var reg = regexp.MustCompile(g.regexp)
	var lineNum int64
	for fileScanner.Scan() {
		lineNum++
		var text = fileScanner.Text()
		if reg.MatchString(text) {
			r := &grepRes{
				Filename: fsrc.Name(),
				Filepath: g.path,
				Number:   lineNum,
				Content:  text,
			}
			res = append(res, r)
		}
	}
	return
}

// 复制一份指针里面的所有参数
func copyGrep(m *Grep) *Grep {
	vt := reflect.TypeOf(m).Elem()
	newoby := reflect.New(vt)
	newoby.Elem().Set(reflect.ValueOf(m).Elem())
	return newoby.Interface().(*Grep)
}
