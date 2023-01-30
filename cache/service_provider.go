package cache

import (
	"github.com/yafgo/framework/cache/console"
	IConsole "github.com/yafgo/framework/contracts/console"
	"github.com/yafgo/framework/facades"
)

type ServiceProvider struct {
}

func (database *ServiceProvider) Register() {
	app := Application{}
	facades.Cache = app.Init()
}

func (database *ServiceProvider) Boot() {
	database.registerCommands()
}

func (database *ServiceProvider) registerCommands() {
	facades.Artisan.Register([]IConsole.Command{
		&console.ClearCommand{},
	})
}
