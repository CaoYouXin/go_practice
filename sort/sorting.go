package main

import (
	"fmt"
	"sort"
)

type byLength []string

func (strs byLength) Len() int {
	return len(strs)
}

func (strs byLength) Swap(i, j int) {
	strs[i], strs[j] = strs[j], strs[i]
}

func (strs byLength) Less(i, j int) bool {
	return len(strs[i]) < len(strs[j])
}

func main() {
	strs := []string{"a", "c", "b"}
	sort.Strings(strs)
	fmt.Println(strs, sort.StringsAreSorted(strs))

	ints := []int{-1, -3, 5, 3}
	sort.Ints(ints)
	fmt.Println(ints, sort.IntsAreSorted(ints))

	byAlphStrs := []string{"weight", "muscie", "fat"}
	sort.Strings(byAlphStrs)
	fmt.Println(byAlphStrs, sort.StringsAreSorted(byAlphStrs))

	byLenStrs := byLength([]string{"weight", "muscie", "fat"})
	sort.Sort(byLenStrs)
	fmt.Println(byLenStrs, sort.IsSorted(byLenStrs))
}
