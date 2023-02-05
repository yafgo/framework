package auth

import (
	"context"

	"github.com/yafgo/framework/auth/access"
	"github.com/yafgo/framework/auth/console"
	contractconsole "github.com/yafgo/framework/contracts/console"
	"github.com/yafgo/framework/facades"
)

type ServiceProvider struct {
}

func (database *ServiceProvider) Register() {
	facades.Auth = NewAuth(facades.Config.GetString("auth.defaults.guard"))
	facades.Gate = access.NewGate(context.Background())
}

func (database *ServiceProvider) Boot() {
	database.registerCommands()
}

func (database *ServiceProvider) registerCommands() {
	facades.Artisan.Register([]contractconsole.Command{
		&console.JwtSecretCommand{},
		&console.PolicyMakeCommand{},
	})
}
