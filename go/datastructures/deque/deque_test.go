package deque

import (
	"aoc/util"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestIntDeque tests must be run in order because the deque is modified by each test.
func TestIntDeque(t *testing.T) {
	tests := []struct {
		name           string
		valToPushFront *int
		valToPushBack  *int
		valToPopFront  *int
		valToPopBack   *int
		wantFront      *int
		wantBack       *int
		wantDeque      []int
		wantErr        error
	}{
		{
			name:      "Empty deque",
			wantDeque: []int{},
		},
		{
			name:           "Push 5 to front of deque",
			valToPushFront: util.IntPtr(5),
			wantFront:      util.IntPtr(5),
			wantBack:       util.IntPtr(5),
			wantDeque:      []int{5},
		},
		{
			name:          "Push 7 to back of deque",
			valToPushBack: util.IntPtr(7),
			wantFront:     util.IntPtr(5),
			wantBack:      util.IntPtr(7),
			wantDeque:     []int{5, 7},
		},
		{
			name:          "Pop front off deque",
			valToPopFront: util.IntPtr(5),
			wantFront:     util.IntPtr(7),
			wantBack:      util.IntPtr(7),
			wantDeque:     []int{7},
		},
		{
			name:          "Pop back off deque",
			valToPopFront: util.IntPtr(7),
			wantDeque:     []int{},
		},
		{
			name:          "Pop back off deque, want err when peeking front",
			valToPopFront: util.IntPtr(7),
			wantDeque:     []int{},
			wantErr:       errors.New("deque is empty"),
		},
		{
			name:         "Pop off empty deque",
			valToPopBack: util.IntPtr(5),
			wantDeque:    []int{},
			wantErr:      errors.New("deque is empty"),
		},
	}
	intDeque := New[int]()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.valToPushFront != nil {
				intDeque.PushFront(*tc.valToPushFront)
			}
			if tc.valToPushBack != nil {
				intDeque.PushBack(*tc.valToPushBack)
			}
			if tc.valToPopFront != nil {
				res, err := intDeque.PopFront()
				if tc.wantErr != nil {
					assert.EqualError(t, err, tc.wantErr.Error())
				} else {
					assert.Equal(t, *tc.valToPopFront, res)
				}
			}
			if tc.valToPopBack != nil {
				res, err := intDeque.PopBack()
				if tc.wantErr != nil {
					assert.EqualError(t, err, tc.wantErr.Error())
				} else {
					assert.Equal(t, *tc.valToPopBack, res)
				}
			}
			if tc.wantFront != nil {
				res, err := intDeque.Front()
				if tc.wantErr != nil {
					assert.EqualError(t, err, tc.wantErr.Error())
				} else {
					assert.Equal(t, *tc.wantFront, res)
				}
			}
			if tc.wantBack != nil {
				res, err := intDeque.Back()
				if tc.wantErr != nil {
					assert.EqualError(t, err, tc.wantErr.Error())
				} else {
					assert.Equal(t, *tc.wantBack, res)
				}
			}
			for i, v := range tc.wantDeque {
				res, err := intDeque.Get(i)
				if tc.wantErr != nil {
					assert.EqualError(t, err, tc.wantErr.Error())
				} else {
					assert.Equal(t, v, res)
				}
			}
		})
	}
}
