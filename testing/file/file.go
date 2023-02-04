package file

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/yafgo/framework/support/file"
)

func GetLineNum(file string) int {
	total := 0
	f, _ := os.OpenFile(file, os.O_RDONLY, 0444)
	buf := bufio.NewReader(f)

	for {
		_, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				total++

				break
			}
		} else {
			total++
		}
	}

	defer func() {
		f.Close()
	}()

	return total
}

func CreateEnv(envPath ...string) (err error) {

	envFile := ".env"
	if len(envPath) > 0 {
		_envPath := strings.TrimSpace(envPath[0])
		if _envPath != "" {
			envFile = _envPath
		}
	}

	content := []byte(`APP_NAME=yafgo
APP_ENV=local
APP_KEY=
APP_DEBUG=true
APP_URL=http://localhost
APP_HOST=0.0.0.0:3000

DB_CONNECTION=mysql
DB_HOST=
DB_PORT=3306
DB_DATABASE=
DB_USERNAME=
DB_PASSWORD=

REDIS_HOST=127.0.0.1
REDIS_PORT=6379
REDIS_PASSWORD=

LOG_CHANNEL=single

MY_PI=3.1415926
MY_TIMEOUT=30s
CACHE_DURATION=2.5h
`)
	err = file.Create(envFile, content)
	return
}
