package file

import (
	"os"
)

func CreateEnv(mode ...string) (err error) {

	envFile := ".env"
	if len(mode) > 0 {
		envFile += "." + mode[0]
	}

	file, err := os.Create(envFile)
	if err != nil {
		return
	}
	defer func() {
		file.Close()
	}()

	_, err = file.WriteString(`APP_NAME=yafgo
APP_ENV=local
APP_KEY=
APP_DEBUG=true
APP_URL=http://localhost
APP_HOST=127.0.0.1:3000

DB_CONNECTION=mysql
DB_HOST=
DB_PORT=3306
DB_DATABASE=
DB_USERNAME=
DB_PASSWORD=

REDIS_HOST=127.0.0.1
REDIS_PASSWORD=
REDIS_PORT=6379

MY_PI=3.1415926
MY_TIMEOUT=30s
CACHE_DURATION=2.5h
`)
	return
}
