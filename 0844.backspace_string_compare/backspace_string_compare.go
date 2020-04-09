package backspace_string_compare

const BACKSPACE = '#'

func reduceBackspaces(input string) string {
	var output []rune
	for _, r := range input {
		if r == BACKSPACE {
			if len(output) > 0 {
				output = output[0 : len(output)-1]
			}
		} else {
			output = append(output, r)
		}
	}
	return string(output)
}

func backspaceCompare(S string, T string) bool {
	return reduceBackspaces(S) == reduceBackspaces(T)
}
