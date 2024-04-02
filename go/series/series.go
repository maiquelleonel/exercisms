package series

func All(n int, s string) []string {
	var ret []string
	for i := 0; i <= len(s)-n; i++ {
		ret = append(ret, s[i:i+n])
	}
	return ret
}

func UnsafeFirst(n int, s string) string {
	return s[0:n]
}

func First(n int, s string) (first string, ok bool) {
	if n <= len(s) {
		ok = true
		first = UnsafeFirst(n, s)
	}
	return
}
