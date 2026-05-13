package pgms

import "fmt"

func generateAllSubarraysV1(arr []int) {
	n := len(arr)

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			subarray := arr[i:j+1]
			fmt.Println(subarray)
		}
	}
}

func generateAllSubarraysV2(arr []int, start, end int) {
	n := len(arr)

	if end == n {
		return
	} else if start > end {
		generateAllSubarraysV2(arr, 0, end+1)
	} else {
		subarray := arr[start:end+1]
		fmt.Println(subarray)
		generateAllSubarraysV2(arr, start+1, end)
	}
}

func Program019() {
	arr := []int{1, 2, 3, 4}
	generateAllSubarraysV2(arr, 0, 0)
}