package tools

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func PrintRed(message string) {
	fmt.Printf("\033[31m%s\033[0m\n", message)
}

func GetTermWidth() int {
	_, w := GetTermDim()
	return w
}

func GetTermDim() (int, int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	termDim, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	var h, w int
	fmt.Sscan(string(termDim), &h, &w)
	return h, w
}

func WipeLine() {
	fmt.Print("\r")
	for i := 0; i < GetTermWidth(); i++ {
		fmt.Print(" ")
	}
	fmt.Print("\r")
}
