package tools

import (
	"testing"
)

func TestTernary(t *testing.T) {
	t.Log("Testing plurals:")
	for i := 0; i < 5; i++ {
		output := Ternary(i > 1, "apples", "apple")
		t.Logf("> %d %v", i, output)
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
  t.Log("Testing booleans:")
  booloutput := Ternary("apple" != "orange", false, true)
  if booloutput == true {
    t.Error("apple is not an orange!")
  } else {
    t.Logf("> apple = orange : %v", booloutput)
  }
  t.Log("Testing integers:")
  intoutput := Ternary(true, 1, 0)
  if intoutput == 1 {
    t.Log("> true is 1")
  } else {
    t.Errorf("true should be 1! however it is: %v", intoutput)
  }
}
