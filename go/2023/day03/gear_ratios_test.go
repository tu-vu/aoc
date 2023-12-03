package day03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGearRatiosPart1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr error
	}{
		{
			name:  "source",
			input: content,
			want:  535351,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := gearRatiosPart1(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestGearRatiosPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "source",
			input: content,
			want:  87287096,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := gearRatiosPart2(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}
