package database

import (
	"context"

	consolecontract "github.com/yafgo/framework/contracts/console"
	"github.com/yafgo/framework/database/console"
	"github.com/yafgo/framework/facades"
)

type ServiceProvider struct {
}

func (database *ServiceProvider) Register() {
	facades.Orm = NewOrm(context.Background())
}

func (database *ServiceProvider) Boot() {
	database.registerCommands()
}

func (database *ServiceProvider) registerCommands() {
	facades.Artisan.Register([]consolecontract.Command{
		&console.MigrateMakeCommand{},
		&console.MigrateCommand{},
		&console.MigrateRollbackCommand{},
		&console.ModelMakeCommand{},
	})
}
