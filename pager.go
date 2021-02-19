/* TODO
   - make it possible to user another pager than less (check $PAGER, default to more)
	 - check the presence of the pager executable in the current $PATH
*/
package tools

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"
)

/*
pipe string to less
(ideally, check that less is installed before running this!)
Source: https://stackoverflow.com/questions/28705716/paging-output-from-go
*/
func Less(message string) error {
	cmd := exec.Command("less", "-FRIX")
	cmd.Stdin = strings.NewReader(message)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	return cmd.Run()
}

/* pipe command to less, using pipe */
func LessPipe(command string, args ...string) (err error) {
	cmd := exec.Command(command, args...)
	less := exec.Command("less", "-FRIX")
	var pipeOut, pipeErr io.ReadCloser
	if pipeOut, err = cmd.StdoutPipe(); err != nil {
		return
	}
	if pipeErr, err = cmd.StderrPipe(); err != nil {
		return
	}
	combinedReader := io.MultiReader(pipeOut, pipeErr)
	less.Stdin = combinedReader // combine outputs of stdout and stderr
	less.Stdout, less.Stderr = os.Stdout, os.Stderr
	if err = less.Start(); err != nil {
		return
	}
	go func() {
		if err = cmd.Run(); err != nil {
			return
		}
	}()
	return less.Wait()
}

/* pipe command to less, simple, reliable, no goroutine, in-memory buffer, slow for
   commands with huge output */
func LessString(command string, args ...string) (err error) {
	var buffer []byte
	if buffer, err = exec.Command(command, args...).CombinedOutput(); err != nil {
		return
	}
	less := exec.Command("less", "-FRIX")
	less.Stdin = bytes.NewReader(buffer)
	less.Stdout, less.Stderr = os.Stdout, os.Stdout
	return less.Run()
}
