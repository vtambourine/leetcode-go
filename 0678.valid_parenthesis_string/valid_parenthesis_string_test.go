package valid_parenthesis_string

import "testing"

type test struct {
	input  string
	expect bool
}

func TestCheckValidString(t *testing.T) {
	tests := []test{
		{"()", true},
		{"(*)", true},
		{"(*))", true},
		{"((*)", true},
		{"(()", false},
		{"())", false},
		{"(*)))", false},
		{"(((*)", false},
		{"(*(()", false},
		{"()()(", false},
		{")(", false},
		{")*(", false},
		{"*(", false},
		{"*(*", true},
		{"((*)(*))((*", false},
		{"(())((())()()(*)(*()(())())())()()((()())((()))(*", false},
	}

	for _, c := range tests {
		if result := checkValidString(c.input); result != c.expect {
			t.Fatalf("checkValidString(%#v) fails.\nExpected %v\nReceived %v", c.input, c.expect, result)
		}
	}
}
