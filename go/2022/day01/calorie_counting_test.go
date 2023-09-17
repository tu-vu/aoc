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
			want:  66487,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := calorieCountingPart1(tc.input)
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
			want:  197301,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := calorieCountingPart2(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}
