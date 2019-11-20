package tools

import (
	"bytes"
	"os"
	"os/exec"
)

/* Passthru execution
   Simplest scenario */
func Xeq(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

/* More complex scenarii, involving options and result structs */

type XeqOptions struct {
	Executable string
	Args       []string
	Stdin      []byte
	Envvars    []string
}

type XeqResult struct {
	XEQ    []string
	Err    error
	Stdout []byte
	Stderr []byte
}

/* No options required */
func XeqR(command string, args ...string) (result XeqResult) {
	options := XeqOptions{command, args, []byte{}, []string{}}
	return XeqOR(options)
}

/* Require options */
func XeqOR(options XeqOptions) (result XeqResult) {
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
