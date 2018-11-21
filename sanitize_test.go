package tools

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestSanitize(t *testing.T) {
	content := "Coucou, ceci est un test !"
	testfile := "/tmp/testfile-sanitize"
	err := ioutil.WriteFile(testfile, []byte(content), 0644)
	if err != nil {
		t.Errorf("Cannot create test file: %q\n", testfile)
	}
	err = Sanitize(testfile, true)
	if err != nil {
		t.Errorf("First pass failed\n")
	}
	err = Sanitize(testfile, true)
	if err != nil {
		t.Errorf("Second pass failed\n")
	}
	os.Remove(testfile)
}
