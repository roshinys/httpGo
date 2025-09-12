package mystrings

func reverse(s string) string {
	res := ""
	for _, v := range s {
		res = string(v) + res
	}
	return res
}
