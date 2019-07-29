package tools

import (
	"fmt"
	"testing"
)

func TestExecute(t *testing.T) {
	err := Execute("date", "+%F %T")
	if err != nil {
		t.Errorf("Execute failed with error: %v\n", err)
	}
}

func TestExecutor(t *testing.T) {
	options := ExecutorOptions{
		Executable: "date",
		Args:       []string{"+%F %T"},
	}
	result := Executor(options)
	displayResult(result)
	options = ExecutorOptions{
		Executable: "nonexistentexecutable",
	}
	result = Executor(options)
	displayResult(result)
	options = ExecutorOptions{
		Executable: "env",
		//Args:       []string{"$FOO", "$BAR"},
		Envvars: []string{"FOO=bar", "BAR=foo"},
	}
	result = Executor(options)
	displayResult(result)
}

func displayResult(result ExecutorResult) {
	fmt.Printf("XEQ: %v\n", result.XEQ)
	fmt.Printf("Err: %v\n", result.Err)
	fmt.Printf("Stdout: %q\n", string(result.Stdout))
	fmt.Printf("Stderr: %q\n", string(result.Stderr))
}
