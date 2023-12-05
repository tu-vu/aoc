package day04

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScratchCardsPart1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    float64
		wantErr error
	}{
		{
			name:  "source",
			input: content,
			want:  26346,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := scratchCardsPart1(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestScratchCardsPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "source",
			input: content,
			want:  8467762,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := scratchCardsPart2(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}
