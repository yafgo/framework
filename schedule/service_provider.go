package schedule

import (
	"github.com/yafgo/framework/facades"
)

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register() {
	facades.Schedule = NewApplication()
}

func (receiver *ServiceProvider) Boot() {

}
