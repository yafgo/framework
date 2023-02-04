package config

import (
	"github.com/yafgo/framework/facades"
)

type ServiceProvider struct {
}

func (config *ServiceProvider) Register() {
	facades.Config = NewApplication()
}

func (config *ServiceProvider) Boot() {

}
