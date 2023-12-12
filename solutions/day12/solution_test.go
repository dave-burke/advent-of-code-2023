package day12

import "testing"

var testLines = map[string]int{
	"???.### 1,1,3":             1,
	".??..??...?##. 1,1,3":      4,
	"?#?#?#?#?#?#?#? 1,3,1,6":   1,
	"????.#...#... 4,1,1":       1,
	"????.######..#####. 1,6,5": 4,
	"?###???????? 3,2,1":        10,
}

func TestCountArrangemen(t *testing.T) {
	for line, expected := range testLines {
		result := CountArrangemen(line)

		if expected != result {
			t.Errorf("%s has %d arrangements, but got %d", line, expected, result)
		}
	}
}
