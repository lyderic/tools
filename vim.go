package tools

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

/*
Edit input string in vim and return edited string
*/
func VimString(input string) (output string, err error) {
	input = input + "\n" // to avoid 'incomplete last line in vim
	var file *os.File
	if file, err = ioutil.TempFile("", "lyderictoolstmp"); err != nil {
		return
	}
	path := file.Name()
	defer os.Remove(path)
	if _, err = file.Write([]byte(input)); err != nil {
		return
	}
	if err = file.Close(); err != nil {
		return
	}
	if err = Vim(path); err != nil {
		return
	}
	var content []byte
	if content, err = ioutil.ReadFile(path); err != nil {
		return
	}
	output = strings.TrimSpace(string(content)) // as vim appends a newline automatically
	return
}

/*
Edit a file path with vim
*/
func Vim(path string) (err error) {
	vim := "vim"
	if _, err = exec.LookPath(vim); err != nil {
		return fmt.Errorf("vim not found on this system!")
	}
	cmd := exec.Command(vim, path)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	if err = cmd.Start(); err != nil {
		return
	}
	if err = cmd.Wait(); err != nil {
		return
	}
	return
}
