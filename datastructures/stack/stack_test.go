package stack

import (
	"aoc/util"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestIntStack tests must be run in order because the stack is modified by each test.
func TestIntStack(t *testing.T) {
	tests := []struct {
		name      string
		valToPush *int
		valToPop  *int
		wantStack Stack[int]
		wantErr   error
	}{
		{
			name:      "empty stack",
			wantStack: Stack[int]{},
		},
		{
			name:      "Push 5 onto stack",
			valToPush: util.IntPtr(5),
			wantStack: Stack[int]{5},
		},
		{
			name:      "Push 7 onto stack",
			valToPush: util.IntPtr(7),
			wantStack: Stack[int]{5, 7},
		},
		{
			name:      "Pop 7 off stack",
			valToPop:  util.IntPtr(7),
			wantStack: Stack[int]{5},
		},
		{
			name:      "Pop 5 off stack",
			valToPop:  util.IntPtr(5),
			wantStack: Stack[int]{},
		},
		{
			name:      "Pop off empty stack",
			valToPop:  util.IntPtr(5),
			wantStack: Stack[int]{},
			wantErr:   errors.New("stack is empty"),
		},
	}
	intStack := Stack[int]{}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.valToPush != nil {
				intStack.Push(*tc.valToPush)
				assert.Equal(t, *tc.valToPush, intStack.Peek())
			}
			if tc.valToPop != nil {
				res, err := intStack.Pop()
				if tc.wantErr != nil {
					assert.EqualError(t, err, tc.wantErr.Error())
				} else {
					assert.Equal(t, *tc.valToPop, res)
				}
			}
			assert.Equal(t, tc.wantStack, intStack)
		})
	}
}
