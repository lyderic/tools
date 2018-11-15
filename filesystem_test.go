package tools

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestPathExists(t *testing.T) {
	fmt.Println("[Filesystem]")
	/* we suppose there will always be a HOME */
	home := os.Getenv("HOME")
	foundHome := PathExists(home)
	if !foundHome {
		t.Error("No HOME found!")
	} else {
		fmt.Printf("> HOME found: %s\n", home)
	}
	imaginaryPath := "/path/to/nonexistent/foobar"
	foundImaginary := PathExists(imaginaryPath)
	if foundImaginary {
		t.Error(imaginaryPath, "should NOT be found!")
	} else {
		fmt.Printf("> Imaginary path NOT found (%s)\n", imaginaryPath)
	}
}

func TestCopy(t *testing.T) {
	src := "a_dummy_file_for_testing_this_api_src"
	dst := "a_dummy_file_for_testing_this_api_dst"
	// Create a dummy file to have something to copy
	err := ioutil.WriteFile(src, []byte("foo bar baz"), 0644)
	if err != nil {
		t.Error("failed to create temporary file")
	}
	err = Copy(src, dst)
	if err != nil {
		t.Error("copy failed")
	} else {
		t.Log("copy successful")
	}
	// check both file are the same!
	srcMd5, _ := Md5(src)
	dstMd5, _ := Md5(dst)
	fmt.Printf("> (copy) MD5 src: %s\n", srcMd5)
	fmt.Printf("> (copy) MD5 dst: %s\n", dstMd5)
	if srcMd5 != dstMd5 {
		t.Error("MD5 checksums don't match")
	}
	os.Remove(src)
	os.Remove(dst)
}

func TestMd5(t *testing.T) {
	file := "a_dummy_file_for_testing_this_api_md5"
	err := ioutil.WriteFile(file, []byte("riri fifi loulou"), 0644)
	if err != nil {
		t.Error("failed to create temporary file")
	}
	md5, err := Md5(file)
	if err != nil {
		t.Error("MD5 computation failed")
	}
	fmt.Printf("> MD5: %v\n", md5)
	if md5 != "d6d37b104b688ceb8be24a0f9598a1df" {
		t.Error("MD5 don't match")
	}
	os.Remove(file)
}
