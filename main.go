package main

import "fmt"

func calculate(s string) int {
	stack := make([]int, 0)
	operand := 0
	res := 0
	sign := 1 // true 1, false -1

	for index, val := range s {
		if val == '\t' || val == '\n' {
			continue
		}
		fmt.Printf("******* CYCLE %v, value: %v ******\n ", index, string(val))
		if isDigit(val) {
			// builds base 10 number
			i, _ := strconv.Atoi(string(val))
			operand = (operand * 10) + i
			//   fmt.Printf("******* operand: %v ******\n ", operand)

		} else if val == '+' {
			//  fmt.Printf("'+' operand: %d sign: %v res: %d \n", operand, sign, res)
			res += sign * operand
			// fmt.Printf("'+' operand: %v sign: %v res: %d \n", operand, sign, res)

			sign = 1
			operand = 0
		} else if val == '-' {
			// fmt.Printf("'-' operand: %d sign: %v res: %d \n", operand, sign, res)
			res += sign * operand
			// fmt.Printf("'-' operand: %d sign: %v res %d \n", operand, sign, res)
			sign = -1
			operand = 0
		} else if val == '(' {
			fmt.Printf("stack.PUSH() res: %d sign: %d \n", res, sign)

			stack = append(stack, res)
			stack = append(stack, sign)
			sign = 1
			res = 0
		} else if val == ')' {
			res += sign * operand

			// fmt.Println("len stack:", len(stack))
			// fmt.Printf("***** res:%v [0]: %v\n", res, stack[len(stack)-1])

			res *= stack[len(stack)-1]
			fmt.Println(res)
			stack = stack[:len(stack)-1]

			//   fmt.Printf("2 ***** res:%v [0]: %v\n", res, stack[len(stack)-1])
			res += stack[len(stack)-1]
			fmt.Println(stack)
			stack = stack[:len(stack)-1]

			operand = 0
		} else {
			continue
		}
	}
	return res + sign*operand
}

func isDigit(s rune) bool {
	if unicode.IsDigit(s) {
		return true
	}
	return false
}
