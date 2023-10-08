package decode_encode

import (
	"github.com/Hamsajj/leetGo/utils"
	"testing"
)

func TestDecodeEncode(t *testing.T) {
	tests := []struct {
		name  string
		input []string
	}{
		{
			name:  "testCase 1",
			input: []string{"lint", "code", "love", "you"},
		},
		{
			name:  "testCase 2",
			input: []string{"we", "say", ":", "yes"},
		},
		{
			name:  "testCase 3",
			input: []string{"1", "2", "3", "123"},
		},
		{
			name:  "testCase 3",
			input: []string{"1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded := encode(tt.input)
			decoded := decode(encoded)
			if !utils.AreSlicesEqual(tt.input, decoded) {
				t.Errorf("decoded (%v) is not the same as input (%v)", decoded, tt.input)
			}
		})
	}
}
