package config

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yafgo/framework/testing/file"
)

func TestInit(t *testing.T) {
	err := file.CreateEnv()
	assert.Nil(t, err)
	assert.NotPanics(t, func() {
		app := Application{}
		app.Init()
	})
}

func TestEnv(t *testing.T) {
	app := Application{}
	app.Init()

	assert.Equal(t, "yafgo", app.GetString("APP_NAME"))
	assert.Equal(t, "127.0.0.1", app.GetString("DB_HOST", "127.0.0.1"))
}

func TestAdd(t *testing.T) {
	app := Application{}
	app.Init()
	app.Add("app", map[string]any{
		"env": "local",
	})

	assert.Equal(t, "local", app.GetString("app.env"))
}

func TestGet(t *testing.T) {
	app := Application{}
	app.Init()

	assert.Equal(t, "yafgo", app.Get("APP_NAME").(string))
}

func TestGetString(t *testing.T) {
	app := Application{}
	app.Init()

	app.Add("database", map[string]any{
		"default": app.Env("DB_CONNECTION", "mysql"),
		"connections": map[string]any{
			"mysql": map[string]any{
				"host": app.Env("DB_HOST", "127.0.0.1"),
			},
		},
	})

	assert.Equal(t, "yafgo", app.GetString("APP_NAME"))
	assert.Equal(t, "127.0.0.1", app.GetString("database.connections.mysql.host"))
	assert.Equal(t, "mysql", app.GetString("database.default"))
}

func TestGetInt(t *testing.T) {
	app := Application{}
	app.Init()

	assert.Equal(t, 3306, app.GetInt("DB_PORT"))
}

func TestGetFloat64(t *testing.T) {
	app := Application{}
	app.Init()

	assert.Equal(t, 3.1415926, app.GetFloat64("MY_PI"))
}

func TestGetDuration(t *testing.T) {
	app := Application{}
	app.Init()

	assert.Equal(t, time.Second*30, app.GetDuration("MY_TIMEOUT"))
	assert.Equal(t, (time.Hour*2)+(time.Minute*30), app.GetDuration("CACHE_DURATION"))
}

func TestGetBool(t *testing.T) {
	app := Application{}
	app.Init()

	assert.Equal(t, true, app.GetBool("APP_DEBUG"))

	err := os.Remove(".env")
	assert.Nil(t, err)
}
