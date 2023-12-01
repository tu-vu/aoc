package day01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalorieCountingPart1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr error
	}{
		{
			name:  "source",
			input: content,
			want:  56465,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := trebutchetPart1(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestCalorieCountingPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "source",
			input: content,
			want:  55902,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := trebutchetPart2(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}
