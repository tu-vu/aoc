package day02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCubeConundrumPart1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr error
	}{
		{
			name:  "source",
			input: content,
			want:  2528,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := cubeConundrumPart1(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestCubeConundrumPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  float64
	}{
		{
			name:  "source",
			input: content,
			want:  67363,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := cubeConundrumPart2(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}
