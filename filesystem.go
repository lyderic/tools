package tools

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"log"
	"os"
	"os/user"
)

func PathExists(path string) (found bool) {
	found = true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		found = false
	}
	return
}

func Copy(from, to string) error {
	if !PathExists(from) {
		return errors.New("tools.Copy: source file not found! Copy aborted.")
	}
	src, err := os.Open(from)
	if err != nil {
		return errors.New("tools.Copy: cannot open source file for reading")
	}
	defer src.Close()
	dst, err := os.Create(to)
	if err != nil {
		return errors.New("tools.Copy: cannot create destination file")
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		return errors.New("tools.Copy: I/O copy failed")
	}
	err = dst.Sync()
	if err != nil {
		return errors.New("tools:Copy: destination failed to sync")
	}
	return nil
}

func GetHomeDir() string {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return user.HomeDir
}

func Md5(path string) (string, error) {
	var output string
	file, err := os.Open(path)
	if err != nil {
		return output, err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return output, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	output = hex.EncodeToString(hashInBytes)
	return output, nil
}
