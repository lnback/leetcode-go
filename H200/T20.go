package H200

func isValid(s string) bool  {
	stack := []string{}
	for _,v := range s{
		ch := string(v)
		if ch == "(" || ch == "[" || ch == "{"{
			stack = append(stack,ch)
		}else {
			if len(stack) == 0{
				return false
			}
			//取栈顶元素
			top := stack[len(stack) - 1]
			//出栈-1
			if top == "(" && ch == ")" || top == "[" && ch == "]" || top == "{" && ch == "}"{
				stack = stack[:len(stack)-1]
			}else {
				return false
			}
		}
	}
	return len(stack) == 0
}
