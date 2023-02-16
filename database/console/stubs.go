package console

type Stubs struct {
}

func (r Stubs) Model() string {
	return `package models

import (
	"github.com/yafgo/framework/database/orm"
)

type DummyModel struct {
	orm.Model
}
`
}
