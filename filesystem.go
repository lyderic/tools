package tools

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

/* return true if a file or directory is found */
func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

/* check if a directory is empty */
func IsDirEmpty(dir string) bool {
	fh, err := os.Open(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return false
	}
	defer fh.Close()
	_, err = fh.Readdir(1)
	if err == io.EOF {
		return true
	}
	return false
}

/* copy a source file to a destination file, overwrite if exists */
func Copy(srcpath, dstpath string) (err error) {
	var src, dst *os.File
	if _, err = os.Stat(srcpath); os.IsNotExist(err) {
		return
	}
	if src, err = os.Open(srcpath); err != nil {
		return
	}
	defer src.Close()
	if dst, err = os.Create(dstpath); err != nil {
		return
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		return
	}
	if err = dst.Sync(); err != nil {
		return
	}
	return
}

/* MD5 hash of a file */
func Md5(path string) (output string, err error) {
	var file *os.File
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()
	hash := md5.New()
	if _, err = io.Copy(hash, file); err != nil {
		return
	}
	hashInBytes := hash.Sum(nil)[:16]
	output = hex.EncodeToString(hashInBytes)
	return
}
