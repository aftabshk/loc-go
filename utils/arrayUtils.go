package utils

func contains(arr []string, s string) (isFound bool) {
	for i := 0; i < len(arr) && !isFound; i++ {
		isFound = arr[i] == s
	}
	return
}
