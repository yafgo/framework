package log

import (
	"context"

	"github.com/yafgo/framework/contracts/log"
)

type Application struct {
}

func (app *Application) Init() log.Log {
	logrusInstance := logrusInstance()

	return NewLogrus(logrusInstance, NewWriter(logrusInstance.WithContext(context.Background())))
}
