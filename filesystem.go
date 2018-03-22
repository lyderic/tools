package tools

import (
	"log"
	"os"
	"os/user"
)

func PathExists(path string) (found bool) {
	found = true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		found = false
	}
	return
}

func GetHomeDir() string {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return user.HomeDir
}
