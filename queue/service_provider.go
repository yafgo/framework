package queue

import (
	"github.com/yafgo/framework/contracts/console"
	"github.com/yafgo/framework/facades"
	queueConsole "github.com/yafgo/framework/queue/console"
)

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register() {
	facades.Queue = NewApplication()
}

func (receiver *ServiceProvider) Boot() {
	receiver.registerCommands()
}

func (receiver *ServiceProvider) registerCommands() {
	facades.Artisan.Register([]console.Command{
		&queueConsole.JobMakeCommand{},
	})
}
