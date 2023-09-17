package day02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRockPaperScissorsPart1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr error
	}{
		{
			name:  "source",
			input: content,
			want:  12586,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := rockPaperScissorsPart1(tc.input)
			if tc.wantErr != nil {
				assert.EqualError(t, err, tc.wantErr.Error())
			} else {
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

func TestRockPaperScissorsPart2(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr error
	}{
		{
			name:  "source",
			input: content,
			want:  13193,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := rockPaperScissorsPart2(tc.input)
			if tc.wantErr != nil {
				assert.EqualError(t, err, tc.wantErr.Error())
			} else {
				assert.Equal(t, tc.want, got)
			}
		})
	}
}
