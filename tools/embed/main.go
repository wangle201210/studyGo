package main

import (
	"embed"
	"fmt"
)

// 如果想加载 . 或者 _ 开头的文件 必须加 /*
// 但是只对当前目录生效，子目录不生效
//
//go:embed static/* static/fe/*
var static embed.FS

func main() {
	s, _ := static.ReadDir("static")
	f, _ := static.ReadDir("static/fe")
	fmt.Printf("s: %+v,f: %+v", s, f)
}
