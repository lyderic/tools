package tools

import (
	"bytes"
	"os"
	"os/exec"
)

func BashExec(script []byte, args ...string) error {
	if len(args) > 0 {
		args = append([]string{"-s"}, args...) // prepending '-s'
	}
	cmd := exec.Command("bash", args...)
	cmd.Stdin = bytes.NewBuffer(script)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	return cmd.Run()
}
