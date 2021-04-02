package repository

import "strings"

type index struct {
	pointers map[string][]int
	items    []string
}

func tokenize(text string) []string {
	arr := strings.Fields(text)
	return convertToLowerCase(arr)
}

func convertToLowerCase(arr []string) []string {
	output := make([]string, len(arr))

	for index, s := range arr {
		output[index] = strings.ToLower(s)
	}

	return output
}

func newIndex(clinicNames []string) index {
	idx := index{
		pointers: map[string][]int{},
		items:    clinicNames,
	}

	idx.fillIndex(clinicNames)

	return idx
}

func (idx *index) fillIndex(clinicNames []string) {
	idx.items = clinicNames

	for clinicIndex, clinicName := range clinicNames {
		tokens := tokenize(clinicName)

		for _, token := range tokens {
			if idx.pointers[token] == nil {
				idx.pointers[token] = []int{}
			}

			idx.pointers[token] = append(idx.pointers[token], clinicIndex)
		}
	}
}

// Search searches the index
func (idx index) Search(text string) []int {
	tokens := tokenize(text)

	result := []int{}

	for _, token := range tokens {
		pointers, ok := idx.pointers[token]
		if !ok {
			continue
		}

		if len(result) == 0 {
			result = pointers
			continue
		}

		result = intersection(result, pointers)
	}

	return result
}

func intersection(a, b []int) []int {
	inA := map[int]bool{}
	inB := map[int]bool{}

	for _, value := range a {
		inA[value] = true
	}

	for _, value := range b {
		inB[value] = true
	}

	inBoth := map[int]bool{}

	for _, value := range a {
		if inA[value] && inB[value] {
			inBoth[value] = true
		}
	}

	for _, value := range b {
		if inA[value] && inB[value] {
			inBoth[value] = true
		}
	}

	return mapKeysToSlice(inBoth)
}

func mapKeysToSlice(m map[int]bool) []int {
	output := make([]int, len(m))

	index := 0
	for key := range m {
		output[index] = key
		index++
	}

	return output
}
