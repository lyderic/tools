package tools

import (
	"fmt"
	"testing"
)

func TestTernary(t *testing.T) {
	fmt.Println("[ternary]")
	for i := 0; i < 5; i++ {
		output := Ternary(i > 1, "apples", "apple")
		fmt.Printf("> %d %v\n", i, output)
		switch i {
		case 0, 1:
			if output != "apple" {
				t.Errorf("ternary not working for %d", i)
			}
		default:
			if output != "apples" {
				t.Errorf("ternary not working for %d", i)
			}
		}
	}
	booloutput := Ternary("apple" != "orange", false, true)
	if booloutput == true {
		t.Error("apple is not an orange!")
	} else {
	fmt.Printf("> apple = orange : %v\n", booloutput)
	}
	intoutput := Ternary(true, 1, 0)
	if intoutput == 1 {
		fmt.Printf("> true is 1\n")
	} else {
		t.Errorf("true should be 1! however it is: %v", intoutput)
	}
}

func TestTimestamp(t *testing.T) {
	fmt.Println("[timestamp]")
	fmt.Println(">", timestamp())
	fmt.Println(">", timestampForFile())
}
