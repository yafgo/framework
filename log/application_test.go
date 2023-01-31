package log

import (
	"os"
	"testing"
	"time"

	"github.com/yafgo/framework/config"
	"github.com/yafgo/framework/facades"
	testingfile "github.com/yafgo/framework/testing/file"

	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	err := testingfile.CreateEnv()
	assert.Nil(t, err)

	addDefaultConfig()

	app := Application{}
	instance := app.Init()

	instance.Debug("debug")
	instance.Error("error")

	dailyFile := "storage/logs/yafgo-" + time.Now().Format("2006-01-02") + ".log"
	singleFile := "storage/logs/yafgo.log"
	singleErrorFile := "storage/logs/yafgo-error.log"

	assert.FileExists(t, dailyFile)
	assert.FileExists(t, singleFile)
	assert.FileExists(t, singleErrorFile)

	assert.Equal(t, 3, testingfile.GetLineNum(dailyFile))
	assert.Equal(t, 3, testingfile.GetLineNum(singleFile))
	assert.Equal(t, 2, testingfile.GetLineNum(singleErrorFile))

	err = os.Remove(".env")
	assert.Nil(t, err)

	err = os.RemoveAll("storage")
	assert.Nil(t, err)
}

// addDefaultConfig Add default config for test.
func addDefaultConfig() {
	configApp := config.ServiceProvider{}
	configApp.Register()

	facadesConfig := facades.Config
	facadesConfig.Add("logging", map[string]interface{}{
		"default": facadesConfig.Env("LOG_CHANNEL", "stack"),
		"channels": map[string]interface{}{
			"stack": map[string]interface{}{
				"driver":   "stack",
				"channels": []string{"daily", "single", "single-error"},
			},
			"single": map[string]interface{}{
				"driver": "single",
				"path":   "storage/logs/yafgo.log",
				"level":  "debug",
			},
			"single-error": map[string]interface{}{
				"driver": "single",
				"path":   "storage/logs/yafgo-error.log",
				"level":  "error",
			},
			"daily": map[string]interface{}{
				"driver": "daily",
				"path":   "storage/logs/yafgo.log",
				"level":  facadesConfig.Env("LOG_LEVEL", "debug"),
				"days":   7,
			},
		},
	})
}
