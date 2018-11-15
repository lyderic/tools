package tools

import (
	"fmt"
	"testing"
)

func TestBashExec(t *testing.T) {
	var err error
	commands := []string{
		"date",
		"date +%F",
		"date | wc -c",
		fmt.Sprintf("echo %q", "hello quoted"),
	}
	for _, command := range commands {
		fmt.Printf("> execution of %q:\n", command)
		if err = BashExec(command); err != nil {
			t.Errorf("command %q failed! %v\n", command, err)
		}
	}
}
