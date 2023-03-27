package files

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

const defaultPath = "./user-data"

var userDataPath, existUserDataPath = os.LookupEnv("CD_USER_DATA_PATH")

func GetPath() string {
	if !existUserDataPath {
		return defaultPath
	}
	return userDataPath
}

func InitPath() error {
	return os.MkdirAll(GetPath(), os.ModePerm)
}

func FileWriter(fh multipart.FileHeader) error {
	file, err := fh.Open()
	defer file.Close()

	if err != nil {
		return err
	}

	dst, err := os.Create(filepath.Join(GetPath(), fh.Filename))
	defer dst.Close()

	_, err = io.Copy(dst, file)
	return err
}
