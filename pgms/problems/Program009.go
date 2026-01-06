package pgms

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumPerimeterOfTriangleV1(arr []int) {
	n := len(arr)
	ans := -1
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			for k := j+1; k < n; k++ {
				if arr[i] + arr[j] > arr[k] && arr[j] + arr[k] > arr[i] && arr[k] + arr[i] > arr[j] {
					ans = max(ans, arr[i] + arr[j] + arr[k])
				}
			}
		}
	}
	fmt.Println(ans)
}

func Program009() {
	arr := []int{6, 1, 6, 5, 8, 4}
	maximumPerimeterOfTriangleV1(arr)
}