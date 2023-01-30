package config

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/spf13/cast"
	"github.com/spf13/viper"

	"github.com/yafgo/framework/contracts/config"
	"github.com/yafgo/framework/support/file"
)

type Application struct {
	vip *viper.Viper
}

var envPrefix = "yafgo"

func (app *Application) Init(mode ...string) config.Config {

	_mode := ""
	if len(mode) > 0 {
		_mode = strings.TrimSpace(mode[0])
	}

	// 默认加载 .env 文件，如果有传参 --env=name 的话，加载 .env.name 文件
	envFile := ".env"
	if _mode != "" {
		envFile = ".env." + _mode
	}

	if !file.Exists(envFile) {
		color.Redln(fmt.Sprintf("Please create %s and initialize it first\nRun command: \ncp .env.example %s", envFile, envFile))
		os.Exit(0)
	}

	app.vip = viper.New()
	app.vip.SetConfigType("env")
	app.vip.AddConfigPath(".")
	app.vip.SetConfigName(envFile)
	err := app.vip.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}

	// 设置环境变量前缀，用以区分 Go 的系统环境变量
	app.vip.SetEnvPrefix(envPrefix)
	// 读取环境变量（支持 flags）
	app.vip.AutomaticEnv()

	return app
}

func SetEnvPrefix(val string) {
	envPrefix = val
}

func (app *Application) Viper() *viper.Viper {
	return app.vip
}

// Env Get config from env.
func (app *Application) Env(envName string, defaultValue ...interface{}) interface{} {
	return app.get(envName, defaultValue...)
}

func (app *Application) get(name string, defaultValue ...interface{}) interface{} {
	if !app.vip.IsSet(name) || empty(app.vip.Get(name)) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return app.vip.Get(name)
}

// Add config to application.
func (app *Application) Add(name string, configuration map[string]interface{}) {
	app.vip.Set(name, configuration)
}

// Get config from application.
//
// @param path eg: "app.name"
func (app *Application) Get(path string, defaultValue ...interface{}) interface{} {
	return app.get(path, defaultValue...)
}

// GetString Get string type config from application.
func (app *Application) GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(app.get(path, defaultValue...))
}

// GetInt Get int type config from application.
func (app *Application) GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(app.get(path, defaultValue...))
}

// GetBool Get bool type config from application.
func (app *Application) GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(app.get(path, defaultValue...))
}

// GetInt64 Get int64 type config from application.
func (app *Application) GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(app.get(path, defaultValue...))
}

// GetUint Get uint type config from application.
func (app *Application) GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(app.get(path, defaultValue...))
}

// GetFloat64 Get float64 type config from application.
func (app *Application) GetFloat64(path string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(app.get(path, defaultValue...))
}

// GetDuration Get time.Duration type config from application.
func (app *Application) GetDuration(path string, defaultValue ...interface{}) time.Duration {
	return cast.ToDuration(app.get(path, defaultValue...))
}

// GetAll Get all config from application.
func (app *Application) GetAll() map[string]interface{} {
	return app.vip.AllSettings()
}

func empty(val interface{}) bool {
	if val == nil {
		return true
	}
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}