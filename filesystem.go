package tools

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func Exists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err // doesn't exist, err != nil
	}
	return nil
}

func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func Copy(from, to string) error {
	var err error
	var src, dst *os.File
	if !PathExists(from) {
		return fmt.Errorf("Copy: source file not found! (%s)", from)
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
	return nil
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
