package config

import (
	"github.com/yafgo/framework/facades"
)

type ServiceProvider struct {
}

func (config *ServiceProvider) Register() {
	app := Application{}
	facades.Config = app.Init()
}

func (config *ServiceProvider) Boot() {

}
