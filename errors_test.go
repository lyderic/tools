package tools

import (
	"fmt"
	"testing"
)

func TestE(t *testing.T) {
	err := fmt.Errorf("Oops there is an error")
	E(err)
}
