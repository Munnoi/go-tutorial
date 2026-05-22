package pgms

import (
	"fmt"
	"sort"
)

func sortByLength(arr []string) {
	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i]) < len(arr[j])
	})
}

func Program037() {
	arr := []string{"GeeksforGeeks", "I", "from", "am"}
	sortByLength(arr)
	fmt.Println(arr)
}