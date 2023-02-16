package cache

import (
	"github.com/yafgo/framework/facades"
)

func prefix() string {
	return facades.Config.GetString("cache.prefix") + ":"
}
