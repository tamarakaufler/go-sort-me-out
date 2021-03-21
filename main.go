package main

import (
	"github.com/tamarakaufler/go-sort-me-out/internal/multi_sort"
	"github.com/tamarakaufler/go-sort-me-out/internal/multiple_key_sort"
	"github.com/tamarakaufler/go-sort-me-out/internal/simple_sort"
)

func main() {
	simple_sort.SortClassByAge(nil)
	simple_sort.SliceSortClassByAge(nil)

	multiple_key_sort.SortClassByMathsResults(nil)
	multiple_key_sort.SortClassByPhysicsResults(nil)
	multiple_key_sort.SortClassByEnglishResults(nil)
	multiple_key_sort.SortClassByChemistryResults(nil)

	multi_sort.SortByAllSubjects(nil)
}
