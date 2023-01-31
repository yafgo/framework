package logger

import (
	"github.com/sirupsen/logrus"
)

type Daily struct {
}

func (daily *Daily) Handle(channel string) (logrus.Hook, error) {
	// todo
	return nil, nil
}
