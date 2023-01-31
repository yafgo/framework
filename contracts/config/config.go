package config

import "time"

//go:generate mockery --name=Config
type Config interface {
	//Env Get config from env.
	Env(envName string, defaultValue ...interface{}) interface{}
	//Add config to application.
	Add(name string, configuration map[string]interface{})
	//Get config from application.
	Get(path string, defaultValue ...interface{}) interface{}
	//Get all config from application.
	GetAll() map[string]interface{}
	//GetString Get string type config from application.
	GetString(path string, defaultValue ...interface{}) string
	//GetInt Get int type config from application.
	GetInt(path string, defaultValue ...interface{}) int
	//GetBool Get bool type config from application.
	GetBool(path string, defaultValue ...interface{}) bool
	//GetInt64 Get int64 type config from application.
	GetInt64(path string, defaultValue ...interface{}) int64
	//GetUint Get uint type config from application.
	GetUint(path string, defaultValue ...interface{}) uint
	//GetFloat64 Get float64 type config from application.
	GetFloat64(path string, defaultValue ...interface{}) float64
	//GetDuration Get time.Duration type config from application.
	GetDuration(path string, defaultValue ...interface{}) time.Duration
}
