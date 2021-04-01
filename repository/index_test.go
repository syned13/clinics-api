package repository

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTokenize(t *testing.T) {
	c := require.New(t)

	tokens := tokenize("Mayo Clinic")
	c.Equal(tokens, []string{"mayo", "clinic"})
}

func TestNewIndex(t *testing.T) {
	c := require.New(t)

	clinicNames := []string{
		"Mayo Clinic",
		"Cleveland Clinic",
		"Mayo Hospital",
	}

	index := NewIndex(clinicNames)
	c.Equal([]int{0, 2}, index.pointers["mayo"])
	c.Equal([]int{0, 1}, index.pointers["clinic"])
}

func TestSearch(t *testing.T) {
	c := require.New(t)

	clinicNames := []string{
		"Mayo Clinic",
		"Cleveland Clinic",
		"Mayo Hospital",
	}

	index := NewIndex(clinicNames)
	results := index.Search("mayo")
	c.Equal([]int{0, 2}, results)
}

func TestIntersection(t *testing.T) {
	c := require.New(t)

	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	inBoth := intersection(a, b)
	sort.Ints(inBoth)
	c.Equal([]int{2, 3}, inBoth)
}
