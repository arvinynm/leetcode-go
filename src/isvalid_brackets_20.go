package main

import "fmt"

func main() {
	bool := isValid("[)]")
	fmt.Println(bool)
}

func isValid(s string) bool {
	bytes := make([]int32, 0)
	if s == "" {
		return true
	}
	if s[0] == ')' || s[0] == ']' || s[0] == '}' {
		return false
	}
	for _, brac := range s {
		switch brac {
		case '(', '[', '{':
			bytes = append(bytes, brac)
		case ')':
			if len(bytes) == 0 {
				return false
			}
			if bytes[len(bytes)-1] == '(' {
				bytes = bytes[:(len(bytes) - 1)]
			} else {
				return false
			}
		case ']':
			if len(bytes) == 0 {
				return false
			}
			if bytes[len(bytes)-1] == '[' {
				bytes = bytes[:(len(bytes) - 1)]
			} else {
				return false
			}
		case '}':
			if len(bytes) == 0 {
				return false
			}
			if bytes[len(bytes)-1] == '{' {
				bytes = bytes[:(len(bytes) - 1)]
			} else {
				return false
			}
		default:
			continue
		}
	}
	if len(bytes) == 0 {
		return true
	}
	return false
}
