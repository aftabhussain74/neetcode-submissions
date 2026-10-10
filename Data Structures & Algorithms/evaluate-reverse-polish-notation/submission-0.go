func evalRPN(tokens []string) int {
	stack := []int{}

	// stack:= append(stack, value)  // push
	// top := stack[len(stack)-1] // top
	// stack:= stack[:len(stack)-1] // pop

for _, token := range tokens {
        switch token {
        case "+":
            b := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            a := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            stack = append(stack, a+b)

        case "-":
            b := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            a := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            stack = append(stack, a-b)

        case "*":
            b := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            a := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            stack = append(stack, a*b)

        case "/":
            b := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            a := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            stack = append(stack, a/b)

        default:
            num, _ := strconv.Atoi(token)
            stack = append(stack, num)
        }
    }

	return stack[len(stack)-1]
}
