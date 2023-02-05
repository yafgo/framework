package console

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yafgo/framework/config"
	"github.com/yafgo/framework/console"
	"github.com/yafgo/framework/contracts"
	console2 "github.com/yafgo/framework/contracts/console"
	"github.com/yafgo/framework/facades"
	"github.com/yafgo/framework/support/file"
	testingfile "github.com/yafgo/framework/testing/file"
)

func TestJobMakeCommand(t *testing.T) {
	err := testingfile.CreateEnv()
	assert.Nil(t, err)

	configApp := config.ServiceProvider{}
	configApp.Register()

	facadesConfig := facades.Config
	facadesConfig.Add("app", map[string]any{
		"providers": []contracts.ServiceProvider{},
	})

	instance := console.NewApplication()
	instance.Register([]console2.Command{
		&JobMakeCommand{},
	})

	assert.NotPanics(t, func() {
		instance.Call("make:job YafgoJob")
	})

	assert.True(t, file.Exists("app/jobs/yafgo_job.go"))
	assert.True(t, file.Remove("app"))
	err = os.Remove(".env")
	assert.Nil(t, err)
}
