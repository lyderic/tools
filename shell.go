package tools

import (
	"bytes"
	"os"
	"os/exec"
)

func BashExec(script string, args ...string) error {
	if len(args) > 0 {
		args = append([]string{"-s"}, args...) // prepending '-s'
	}
	cmd := exec.Command("bash", args...)
	cmd.Stdin = bytes.NewBufferString(script)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	return cmd.Run()
}

func BashExecCapture(script string, args ...string) (stdout, stderr []byte, err error) {
	if len(args) > 0 {
		args = append([]string{"-s"}, args...) // prepending '-s'
	}
	cmd := exec.Command("bash", args...)
	cmd.Stdin = bytes.NewBufferString(script)
	var outBuffer, errBuffer bytes.Buffer
	cmd.Stdout, cmd.Stderr = &outBuffer, &errBuffer
	err = cmd.Run()
	stdout, stderr = outBuffer.Bytes(), errBuffer.Bytes()
	return
}
