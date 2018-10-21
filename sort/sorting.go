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

	strsByLen := []string{"annie", "joe", "youxin"}
	byLenStrs := byLength(strsByLen)
	sort.Sort(byLenStrs)
	fmt.Println(strsByLen, byLenStrs, sort.IsSorted(byLenStrs), sort.StringsAreSorted(strsByLen))
}
