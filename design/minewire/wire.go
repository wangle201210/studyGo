//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package minewire

import (
	"github.com/google/wire"
	"github.com/wangle201210/studyGo/design/minewire/lib"
	"github.com/wangle201210/studyGo/design/minewire/lib1"
)

// wireApp init kratos application.
func startApp() *App {
	panic(wire.Build(
		lib.New,
		lib1.ProviderSet,
		New,
	))
}
