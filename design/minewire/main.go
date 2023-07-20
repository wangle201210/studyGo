package minewire

import (
	"github.com/wangle201210/studyGo/design/minewire/lib"
	"github.com/wangle201210/studyGo/design/minewire/lib1"
)

type App struct {
}

func New(lib *lib.Lib, lib1 *lib1.Lib1) *App {
	return &App{}
}

// 使用wire生成的方法
// 最后生成的 startApp 实际上和 simpleNew 的写法是一样的
func wireNew() *App {
	return startApp()
}

// 不使用wire的生成方法
func simpleNew() *App {
	l := lib.New()
	l1 := lib1.New(l)
	return New(l, l1)
}
