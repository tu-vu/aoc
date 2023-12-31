package util

func IntPtr(i int) *int {
	return &i
}

func StringPtr(s string) *string {
	return &s
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
