package tools

import (
	"fmt"
	"testing"
)

func TestThousandSeparator(t *testing.T) {
	var samples = map[int]string{
		1:          "1",
		42:         "42",
		1970:       "1,970",
		192898:     "192,898",
		190777890:  "190,777,890",
		1000000000: "1,000,000,000",
		999:        "999",
	}
	for k, v := range samples {
		fmt.Printf("> %v %v\n", k, v)
		if ThousandSeparator(k) != v {
			t.Errorf("ThousandSeparator doesn't work for %d [%s]!", k, v)
		}
	}
}

func TestColors(t *testing.T) {
	PrintRedln("Printredln")
	PrintYellow("YELLOW")
	PrintRed("RED")
	PrintColorln(BLUE, "BLUE")
	PrintColorln(CYAN, "Printcolorln")
	PrintGreenf("PrintGreenf : %03d\n", 42)
	for i := 30; i < 38; i++ {
		PrintColor(i, "GO")
	}
	fmt.Println()
}
