

package fileutil

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/astaxie/beego/logs"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func Exists(file string) bool {
	_, err := os.Stat(file)
	return !(err != nil && os.IsNotExist(err))
}

func TryRemoveFile(file string) {
	os.Remove(file)
}

func SetExecutable(file string) error {
	fileInfo, err := os.Stat(file)
	if err != nil {
		return err
	}
	return os.Chmod(file, fileInfo.Mode()|0111)
}

func GetFileMd5(file string) (string, error) {
	if !Exists(file) {
		return "", nil
	}

	pFile, err := os.Open(file)
	defer pFile.Close()
	if err != nil {
		return "", err
	}

	md5h := md5.New()
	io.Copy(md5h, pFile)
	return hex.EncodeToString(md5h.Sum(nil)), nil
}

func CopyFile(src string, dst string, overwrite bool) (written int64, err error) {
	logs.Info("copy file from : " + src + ", to: " + dst)

	srcStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}
	if srcStat.IsDir() {
		return 0, errors.New("src is a directory")
	}

	dstStat, err := os.Stat(dst)
	if !(err != nil && os.IsNotExist(err)) {
		if !overwrite {
			return 0, errors.New("dst file exists")
		}
		if dstStat.IsDir() {
			return 0, errors.New("dst is a directory")
		}
		if err = os.Remove(dst); err != nil {
			return 0, err
		}
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer dstFile.Close()

	types, err := io.Copy(dstFile, srcFile)
	return types, err
}

func GetString(file string) (string, error) {
	fileStr, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(fileStr), nil
}

func GetPid(file string) (int, error) {
	pidStr, err := GetString(file)
	if err != nil && !os.IsNotExist(err) {
		return 0, err
	}

	pid, err := strconv.Atoi(pidStr)
	return pid, nil
}

func WriteString(file, str string) error {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if os.IsNotExist(err) {
		f, err = os.Create(file)
	}
	if err != nil {
		return err
	}

	if err = f.Truncate(0); err != nil {
		return err
	}

	if _, err = f.WriteString(str); err != nil {
		return err
	}

	return nil
}
