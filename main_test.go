package main

import (
	"testing"

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
		{"multiple", []int{1, 3, 2, 5, 4}, 5},
		{"negative", []int{-1, -3, -2, -5, -4}, -1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := maximum(test.data)
			require.Equal(t, test.want, got)
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
		{"multiple", []int{1, 3, 2, 5, 4}, 5},
		{"negative", []int{-1, -3, -2, -5, -4}, -1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := maxChunks(test.data)
			require.Equal(t, test.want, got)
		})
	}
}
