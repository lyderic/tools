package tools

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func Copy(from, to string) (err error) {
	var src, dst *os.File
	if _,err = os.Stat(from);os.IsNotExist(err) {
		return err
	}
	if src, err = os.Open(from); err != nil {
		return err
	}
	defer src.Close()
	if dst, err = os.Create(to); err != nil {
		return err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	if err = dst.Sync(); err != nil {
		return err
	}
	return err
}

func Md5(path string) (string, error) {
	var err error
	var file *os.File
	var output string
	if file, err = os.Open(path); err != nil {
		return output, err
	}
	defer file.Close()
	hash := md5.New()
	if _, err = io.Copy(hash, file); err != nil {
		return output, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	output = hex.EncodeToString(hashInBytes)
	return output, nil
}
