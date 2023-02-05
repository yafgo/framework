package database

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/yafgo/framework/contracts/database"
	"github.com/yafgo/framework/contracts/database/orm"
	"github.com/yafgo/framework/database/support"
	"github.com/yafgo/framework/facades"
)

type DB struct {
	connection      string
	defaultInstance database.Sqlx
	instances       map[string]database.Sqlx
}

func (r *DB) Connection(name string) database.DB {
	defaultConnection := facades.Config.GetString("database.default")
	if name == "" {
		name = defaultConnection
	}

	r.connection = name
	if r.instances == nil {
		r.instances = make(map[string]database.Sqlx)
	}

	if _, exist := r.instances[name]; exist {
		return r
	}

	dsn, err := GetDsn(name)
	if err != nil {
		facades.Log.Errorf("get database dsn error: %v", err)

		return r
	}

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		facades.Log.Errorf("orm open database error: %v", err)

		return r
	}

	r.instances[name] = db

	if name == defaultConnection {
		r.defaultInstance = db
	}

	return r
}

func (r *DB) Query() database.Sqlx {
	if r.connection == "" {
		if r.defaultInstance == nil {
			r.Connection("")
		}

		return r.defaultInstance
	}

	instance, exist := r.instances[r.connection]
	if !exist {
		return nil
	}

	r.connection = ""

	return instance
}

func (r *DB) Transaction(ctx context.Context, txFunc func(tx *sqlx.Tx) error) error {
	tx, err := r.Query().BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	if err := txFunc(tx); err != nil {
		if err := tx.Rollback(); err != nil {
			return errors.Wrapf(err, "rollback error: %v", err)
		}

		return err
	} else {
		return tx.Commit()
	}
}

func GetDsn(connection string) (string, error) {
	driver := facades.Config.GetString("database.connections." + connection + ".driver")
	switch orm.Driver(driver) {
	case orm.DriverMysql:
		return support.GetMysqlDsn(connection), nil
	case orm.DriverPostgresql:
		return support.GetPostgresqlDsn(connection), nil
	case orm.DriverSqlite:
		return support.GetSqliteDsn(connection), nil
	case orm.DriverSqlserver:
		return support.GetSqlserverDsn(connection), nil
	default:
		return "", errors.New("database driver only support mysql, postgresql, sqlite and sqlserver")
	}
}
