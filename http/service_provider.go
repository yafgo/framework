package http

import (
	consolecontract "github.com/yafgo/framework/contracts/console"
	"github.com/yafgo/framework/facades"
	"github.com/yafgo/framework/http/console"
)

type ServiceProvider struct {
}

func (database *ServiceProvider) Register() {
}

func (database *ServiceProvider) Boot() {
	database.registerCommands()
}

func (database *ServiceProvider) registerCommands() {
	facades.Artisan.Register([]consolecontract.Command{
		&console.RequestMakeCommand{},
	})
}
