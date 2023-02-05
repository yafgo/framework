package log

import (
	"github.com/yafgo/framework/contracts/log"
)

type Application struct {
}

func NewApplication(writer log.Writer) log.Log {
	return &Logrus{
		Writer: writer,
	}
}
