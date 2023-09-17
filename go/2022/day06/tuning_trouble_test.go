package day06

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTuningTablePart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "test 1",
			input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want:  5,
		},
		{
			name:  "test 2",
			input: "nppdvjthqldpwncqszvftbrmjlhg",
			want:  6,
		},
		{
			name:  "test 3",
			input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want:  10,
		},
		{
			name:  "test 4",
			input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want:  11,
		},
		{
			name:  "test 5",
			input: content,
			want:  1953,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tuningTables(tc.input, 1)
			if assert.NoError(t, err) {
				assert.Equal(t, got, tc.want)
			}
		})
	}
}

func TestTuningTablePart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "test 1",
			input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			want:  19,
		},
		{
			name:  "test 2",
			input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want:  23,
		},
		{
			name:  "test 3",
			input: "nppdvjthqldpwncqszvftbrmjlhg",
			want:  23,
		},
		{
			name:  "test 4",
			input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want:  29,
		},
		{
			name:  "test 5",
			input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want:  26,
		},
		{
			name:  "test 6",
			input: content,
			want:  2301,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tuningTables(tc.input, 2)
			if assert.NoError(t, err) {
				assert.Equal(t, got, tc.want)
			}
		})
	}
}
