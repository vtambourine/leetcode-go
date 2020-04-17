package valid_parenthesis_string

func checkValidSubstring(s *[]byte, i, n int) bool {
	if len(*s) == i {
		return n == 0
	}
	if n < 0 {
		return false
	}
	if (*s)[i] == '(' {
		return checkValidSubstring(s, i+1, n+1)
	} else if (*s)[i] == ')' {
		return checkValidSubstring(s, i+1, n-1)
	} else {
		return checkValidSubstring(s, i+1, n+1) || checkValidSubstring(s, i+1, n-1) || checkValidSubstring(s, i+1, n)
	}
}

func checkValidString(s string) bool {
	b := []byte(s)
	return checkValidSubstring(&b, 0, 0)
}
