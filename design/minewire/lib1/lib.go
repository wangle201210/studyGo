package lib1

import (
	"github.com/google/wire"
	"github.com/wangle201210/studyGo/design/minewire/lib"
)

var ProviderSet = wire.NewSet(New)

type Lib1 struct {
}

func New(data *lib.Lib) *Lib1 {
	return &Lib1{}
}
