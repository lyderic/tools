package tools

import (
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
		t.Logf("%v %v", k, v)
		if ThousandSeparator(k) != v {
			t.Errorf("ThousandSeparator doesn't work for %d [%s]!", k, v)
		}
	}
}
