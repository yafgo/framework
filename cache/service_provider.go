package cache

import (
	"github.com/yafgo/framework/cache/console"
	i_console "github.com/yafgo/framework/contracts/console"
	"github.com/yafgo/framework/facades"
)

type ServiceProvider struct {
}

func (database *ServiceProvider) Register() {
	facades.Cache = NewApplication()
}

func (database *ServiceProvider) Boot() {
	database.registerCommands()
}

func (database *ServiceProvider) registerCommands() {
	facades.Artisan.Register([]i_console.Command{
		&console.ClearCommand{},
	})
}
