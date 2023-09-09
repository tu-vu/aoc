package day03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRucksackReorganizationPart1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr error
	}{
		{
			name:  "source",
			input: content,
			want:  8139,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := rucksackReorganizationPart1(tc.input)
			if tc.wantErr != nil {
				assert.EqualError(t, err, tc.wantErr.Error())
			} else {
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

func TestRucksackReorganizationPart2(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr error
	}{
		{
			name:  "source",
			input: content,
			want:  2668,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := rucksackReorganizationPart2(tc.input)
			if tc.wantErr != nil {
				assert.EqualError(t, err, tc.wantErr.Error())
			} else {
				assert.Equal(t, tc.want, got)
			}
		})
	}
}
