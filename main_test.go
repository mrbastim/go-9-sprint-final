package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{"zero", 0},
		{"small", 10},
		{"large", 10_000},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := generateRandomElements(test.size)
			require.Equal(t, test.size, len(got))
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty", []int{}, 0},
		{"single", []int{42}, 42},
		{"multiple5", []int{1, 3, 2, 5, 4}, 5},
		{"multiple9", []int{1, 3, 2, 5, 4, 6, 7, 8, 9}, 9},
		{"multiple11", []int{7, 2, 9, 3, 6, 4, 10, 11, 5, 9}, 11},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := maximum(test.data)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty", []int{}, 0},
		{"single", []int{42}, 42},
		{"multiple5", []int{1, 3, 2, 5, 4}, 5},
		{"multiple9", []int{1, 3, 2, 5, 4, 6, 7, 8, 9}, 9},
		{"multiple11", []int{7, 2, 9, 3, 6, 4, 10, 11, 5, 9}, 11},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := maxChunks(test.data)
			assert.Equal(t, test.want, got)
		})
	}
}
