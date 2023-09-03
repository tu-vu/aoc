package day01

import (
	"testing"
)

func TestCalorieCountingPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "source",
			input: content,
			want:  66487,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := calorieCountingPart1(tc.input); got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
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
			if got := calorieCountingPart2(tc.input); got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
