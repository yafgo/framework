package filesystem

import (
	"bytes"
	"io"
	"mime/multipart"
	"path"
	"strings"

	"github.com/yafgo/framework/contracts/filesystem"
	"github.com/yafgo/framework/support/file"
)

const MaxFileNum = 1000

func fileHeaderToString(fileHeader *multipart.FileHeader) (string, error) {
	src, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func fullPathOfFile(filePath string, source filesystem.File, name string) (string, error) {
	extension := path.Ext(name)
	if extension == "" {
		extension = file.Extension(source.File())
	} else {
		extension = strings.TrimLeft(extension, ".")
	}
	return path.Join(filePath, path.Base(name)) + "." + extension, nil
}

func validPath(path string) string {
	realPath := strings.TrimPrefix(path, "./")
	realPath = strings.TrimPrefix(realPath, "/")
	realPath = strings.TrimPrefix(realPath, ".")
	if realPath != "" && !strings.HasSuffix(realPath, "/") {
		realPath += "/"
	}

	return realPath
}
