package tools

import (
	"fmt"
	"testing"
)

func TestXeq(t *testing.T) {
	err := Xeq("date", "+%F %T")
	if err != nil {
		t.Errorf("Execute failed with error: %v\n", err)
	}
}

func TestXeqOR(t *testing.T) {
	options := XeqOptions{
		Executable: "date",
		Args:       []string{"+%F %T"},
	}
	result := XeqOR(options)
	displayResult(result)
	options = XeqOptions{
		Executable: "nonexistentexecutable",
	}
	result = XeqOR(options)
	displayResult(result)
	options = XeqOptions{
		Executable: "env",
		//Args:       []string{"$FOO", "$BAR"},
		Envvars: []string{"FOO=bar", "BAR=foo"},
	}
	result = XeqOR(options)
	displayResult(result)
}

func displayResult(result XeqResult) {
	fmt.Printf("XEQ: %v\n", result.XEQ)
	fmt.Printf("Err: %v\n", result.Err)
	fmt.Printf("Stdout: %q\n", string(result.Stdout))
	fmt.Printf("Stderr: %q\n", string(result.Stderr))
}
