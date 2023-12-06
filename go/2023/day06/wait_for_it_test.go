package day06

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWaitForItPart1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    float64
		wantErr error
	}{
		{
			name:  "source",
			input: content,
			want:  5133600,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := waitForItPart1(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestWaitForItPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  float64
	}{
		{
			name:  "source",
			input: content,
			want:  40651271,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := waitForItPart2(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}
