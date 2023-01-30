package console

import (
	"github.com/yafgo/framework/contracts/console"
)

type Application struct {
}

func (app *Application) Init() console.Artisan {
	return NewCli()
}
