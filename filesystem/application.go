package filesystem

import (
	"context"
	"fmt"

	"github.com/yafgo/framework/contracts/filesystem"
	"github.com/yafgo/framework/facades"
)

type Driver string

const (
	DriverLocal  Driver = "local"
	DriverS3     Driver = "s3"
	DriverOss    Driver = "oss"
	DriverCos    Driver = "cos"
	DriverCustom Driver = "custom"
)

type Storage struct {
	filesystem.Driver
	drivers map[string]filesystem.Driver
}

func NewStorage() *Storage {
	defaultDisk := facades.Config.GetString("filesystems.default")
	if defaultDisk == "" {
		facades.Log.Errorf("[filesystem] please set default disk")
		return nil
	}

	driver, err := NewDriver(defaultDisk)
	if err != nil {
		facades.Log.Errorf("[filesystem] init %s disk error: %+v", defaultDisk, err)
		return nil
	}

	drivers := make(map[string]filesystem.Driver)
	drivers[defaultDisk] = driver
	return &Storage{
		Driver:  driver,
		drivers: drivers,
	}
}

func NewDriver(disk string) (filesystem.Driver, error) {
	ctx := context.Background()
	driver := Driver(facades.Config.GetString(fmt.Sprintf("filesystems.disks.%s.driver", disk)))
	switch driver {
	case DriverLocal:
		return NewLocal(disk)
	case DriverOss:
		return NewOss(ctx, disk)
	case DriverCos:
		return NewCos(ctx, disk)
	case DriverS3:
		return NewS3(ctx, disk)
	case DriverCustom:
		driver, ok := facades.Config.Get(fmt.Sprintf("filesystems.disks.%s.via", disk)).(filesystem.Driver)
		if !ok {
			return nil, fmt.Errorf("[filesystem] init %s disk fail: via must be filesystem.Driver.", disk)
		}

		return driver, nil
	}

	return nil, fmt.Errorf("[filesystem] invalid driver: %s, only support local, s3, oss, cos, custom.", driver)
}

func (r *Storage) Disk(disk string) filesystem.Driver {
	if r.drivers[disk] != nil {
		return r.drivers[disk]
	}

	driver, err := NewDriver(disk)
	if err != nil {
		facades.Log.Error(err.Error())
	}

	return driver
}
