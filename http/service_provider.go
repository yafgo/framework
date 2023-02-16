package http

import (
	consolecontract "github.com/yafgo/framework/contracts/console"
	"github.com/yafgo/framework/facades"
	"github.com/yafgo/framework/http/console"
)

type ServiceProvider struct {
}

func (sp *ServiceProvider) Register() {
}

func (sp *ServiceProvider) Boot() {
	sp.registerCommands()
}

func (sp *ServiceProvider) registerCommands() {
	facades.Artisan.Register([]consolecontract.Command{
		&console.RequestMakeCommand{},
		&console.ControllerMakeCommand{},
		&console.MiddlewareMakeCommand{},
	})
}
