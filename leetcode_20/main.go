package leetcode20

import "strings"

func isValid(s string) bool {
	splitText := strings.Split(s, "")

	if len(splitText)%2 != 0 {
		return false
	}

	for i := 0; i < len(splitText); i += 2 {
		switch splitText[i] {
		case "{":
			if splitText[i+1] != "}" {
				return false
			}
		case "[":
			if splitText[i+1] != "]" {
				return false
			}
		case "(":
			if splitText[i+1] != ")" {
				return false
			}
		}
	}
	return true
}
