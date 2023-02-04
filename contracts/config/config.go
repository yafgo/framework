package config

import "time"

//go:generate mockery --name=Config
type Config interface {
	//Env Get config from env.
	Env(envName string, defaultValue ...any) any
	//Add config to application.
	Add(name string, configuration map[string]any)
	//Get config from application.
	Get(path string, defaultValue ...any) any
	//Get all config from application.
	GetAll() map[string]any
	//GetString Get string type config from application.
	GetString(path string, defaultValue ...any) string
	//GetInt Get int type config from application.
	GetInt(path string, defaultValue ...any) int
	//GetBool Get bool type config from application.
	GetBool(path string, defaultValue ...any) bool
	//GetInt64 Get int64 type config from application.
	GetInt64(path string, defaultValue ...any) int64
	//GetUint Get uint type config from application.
	GetUint(path string, defaultValue ...any) uint
	//GetFloat64 Get float64 type config from application.
	GetFloat64(path string, defaultValue ...any) float64
	//GetDuration Get time.Duration type config from application.
	GetDuration(path string, defaultValue ...any) time.Duration
}
