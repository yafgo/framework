package mail

import (
	"github.com/yafgo/framework/contracts/queue"
	"github.com/yafgo/framework/facades"
)

type ServiceProvider struct {
}

func (route *ServiceProvider) Register() {
	facades.Mail = NewApplication()
}

func (route *ServiceProvider) Boot() {
	facades.Queue.Register([]queue.Job{
		&SendMailJob{},
	})
}
