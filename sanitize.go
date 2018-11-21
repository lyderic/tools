package tools

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Sanitize(path string, verbose bool) (err error) {
	start := time.Now()
	var nbytes int
	if _, err = os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("cannot sanitize %s: file not found", path)
	}
	if nbytes, err = substitute(path,
		"---", "***",
		"--", "—",
		"...", "…",
		"« ", "« ",
		" »", " »",
		" ?", " ?",
		" !", " !",
		" ;", " ;",
		" :", " :",
		" \n", "\n",
		"  ", " "); err != nil {
		return err
	}
	if verbose {
		fmt.Printf("%s: processed %d bytes in %s\n", filepath.Base(path), nbytes, time.Since(start))
	}
	return
}

func substitute(path string, replacements ...string) (n int, err error) {
	r := strings.NewReplacer(replacements...)
	// get content of the file
	var content []byte
	if content, err = ioutil.ReadFile(path); err != nil {
		return n, err
	}
	// create file to write to
	var f *os.File
	if f, err = os.Create(path); err != nil {
		return n, err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	if n, err = r.WriteString(w, string(content)); err != nil {
		return n, err
	}
	w.Flush()
	return
}
