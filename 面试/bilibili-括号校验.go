package main

func isValid(s string) bool {
	hashmap := map[string]string{")":"(", "}":"{", "]":"["}
	stack := make([]string, 0)
	if s == ""{
		return true
	}
	for i := 0; i < len(s); i++{
		if s[i] == '(' || s[i] == '{' || s[i] == '['{
			stack = append(stack, string(s[i]))
		}else if len(stack) > 0 && stack[len(stack) - 1] == hashmap[string(s[i])]{
			stack = stack[:len(stack) - 1]
		}else{
			return false
		}
	}
	return len(stack) == 0
}
