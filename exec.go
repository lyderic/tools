package tools

import (
	"bytes"
	"os"
	"os/exec"
)

func Execute(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

type ExecutorOptions struct {
	Executable string
	Args       []string
	Stdin      []byte
	Envvars    []string
}

type ExecutorResult struct {
	XEQ    []string
	Err    error
	Stdout []byte
	Stderr []byte
}

func Executor(options ExecutorOptions) (result ExecutorResult) {
	cmd := exec.Command(options.Executable, options.Args...)
	result.XEQ = cmd.Args
	if len(options.Stdin) != 0 {
		cmd.Stdin = os.Stdin
	} else {
		cmd.Stdin = bytes.NewBuffer(options.Stdin)
	}
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Env = append(os.Environ(), options.Envvars...)
	err := cmd.Run()
	result.Err = err
	result.Stdout = stdout.Bytes()
	result.Stderr = stderr.Bytes()
	return
}
