package day04

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCampCleanupPart1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr error
	}{
		{
			name:  "source",
			input: content,
			want:  528,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := campCleanupPart1(tc.input)
			if tc.wantErr != nil {
				assert.EqualError(t, err, tc.wantErr.Error())
			} else {
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

func TestCampCleanupPart2(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr error
	}{
		{
			name:  "source",
			input: content,
			want:  881,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := campCleanupPart2(tc.input)
			if tc.wantErr != nil {
				assert.EqualError(t, err, tc.wantErr.Error())
			} else {
				assert.Equal(t, tc.want, got)
			}
		})
	}
}
