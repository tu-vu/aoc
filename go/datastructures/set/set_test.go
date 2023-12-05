package set

import (
	"aoc/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestIntSet tests must be run in order because the set is modified by each test.
func TestIntSet(t *testing.T) {
	tests := []struct {
		name        string
		valToAdd    *int
		valToRemove *int
		wantSet     *Set[int]
		wantErr     error
	}{
		{
			name: "empty set",
			wantSet: &Set[int]{
				set: map[int]struct{}{},
			},
		},
		{
			name:     "Add 5 onto set",
			valToAdd: util.IntPtr(5),
			wantSet: &Set[int]{
				set: map[int]struct{}{
					5: {},
				},
			},
		},
		{
			name:     "Add 7 onto set",
			valToAdd: util.IntPtr(7),
			wantSet: &Set[int]{
				set: map[int]struct{}{
					5: {},
					7: {},
				},
			},
		},
		{
			name:     "Add 5 onto set",
			valToAdd: util.IntPtr(5),
			wantSet: &Set[int]{
				set: map[int]struct{}{
					5: {},
					7: {},
				},
			},
		},
		{
			name:        "Remove 5 off set",
			valToRemove: util.IntPtr(5),
			wantSet: &Set[int]{
				set: map[int]struct{}{
					7: {},
				},
			},
		},
	}
	intSet := New[int]()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.valToAdd != nil {
				intSet.Add(*tc.valToAdd)
			}
			if tc.valToRemove != nil {
				intSet.Remove(*tc.valToRemove)
			}
			assert.Equal(t, tc.wantSet, intSet)
		})
	}
}
