package main

func isValid(s string) bool {
	var stack []rune
	var validStatus bool = true
	if len(s) == 1 {
		return false
	}
	for _, value := range s {
		if '(' == value || '[' == value || '{' == value {
			stack = append(stack, value)
		} else {
			if len(stack) == 0 {
				validStatus = false
				break
			}
			switch value {
			case ')':
				{
					top := stack[len(stack)-1]
					if top != '(' {
						validStatus = false
						break
					}
					stack = stack[:len(stack)-1]
				}
			case ']':
				top := stack[len(stack)-1]
				if top != '[' {
					validStatus = false
					break
				}
				stack = stack[:len(stack)-1]
			case '}':
				top := stack[len(stack)-1]
				if top != '{' {
					validStatus = false
					break
				}
				stack = stack[:len(stack)-1]
			default:
				//
			}
		}
	}
	if len(stack) != 0 {
		validStatus = false
	}
	return validStatus
}
