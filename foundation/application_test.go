package foundation

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yafgo/framework/config"
	"github.com/yafgo/framework/console"
	"github.com/yafgo/framework/contracts"
	"github.com/yafgo/framework/facades"
)

func TestInit(t *testing.T) {
	configApp := config.ServiceProvider{}
	configApp.Register()

	facadesConfig := facades.Config
	facadesConfig.Add("app", map[string]interface{}{
		"providers": []contracts.ServiceProvider{
			&console.ServiceProvider{},
		},
	})

	assert.NotPanics(t, func() {
		app := Application{}
		app.Boot()
	})
}
