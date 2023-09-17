package day05

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSupplyStacksPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "source",
			input: content,
			want:  "TQRFCBSJJ",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := supplyStacks(tc.input, 1)
			if assert.NoError(t, err) {
				assert.Equal(t, got, tc.want)
			}
		})
	}
}

func TestSupplyStacksPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "source",
			input: content,
			want:  "RMHFJNVFP",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := supplyStacks(tc.input, 2)
			if assert.NoError(t, err) {
				assert.Equal(t, got, tc.want)
			}
		})
	}
}
