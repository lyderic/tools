package tools

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

func BashExec(command string) (err error) {
	var file *os.File
	if file, err = ioutil.TempFile("",
		"tools.lyderic.com.script."); err != nil {
		return
	}
	file.Chmod(0700)
	defer os.Remove(file.Name())
	script := fmt.Sprintf("#!/bin/bash\n%s\n", command)
	io.WriteString(file, script)
	file.Close()
	cmd := exec.Command(file.Name())
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	return cmd.Run()
}
