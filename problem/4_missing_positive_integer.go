package problem

func MissingPositiveInt(arr []int) int {
	for i := 0; i < len(arr); {
		if arr[i] < 1 || arr[i] >= len(arr) || arr[i] == i+1 {
			i++
			continue
		}
		arr[arr[i]-1], arr[i] = arr[i], arr[arr[i]-1]
	}

	for i := range arr {
		if arr[i] != i+1 {
			return i + 1
		}
	}

	return len(arr) + 1
}
