package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	t.Run("testing fibonacci func", func(t *testing.T) {

		tests := []struct {
			arg  int
			want int
		}{
			{
				arg:  1,
				want: 1,
			},
			{
				arg:  2,
				want: 1,
			},
			{
				arg:  3,
				want: 2,
			},
		}

		fmt.Println()

		for _, tt := range tests {
			assert.Equal(t, tt.want, fibonacci(tt.arg))
		}
	})
}
