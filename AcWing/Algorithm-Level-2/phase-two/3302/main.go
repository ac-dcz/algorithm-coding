package main

import "fmt"

func Cacl(expr string) int {
	opnum, opand := make([]int, 0), make([]byte, 0)
	expr = "(" + expr + ")"
	for i := 0; i < len(expr); i++ {
		if '0' <= expr[i] && expr[i] <= '9' {
			num := 0
			for '0' <= expr[i] && expr[i] <= '9' {
				num = num*10 + int(expr[i]-'0')
				i++
			}
			opnum = append(opnum, num)
			i--
		} else if expr[i] == '(' {
			opand = append(opand, '(')
		} else {
			for len(opand) > 0 {
				op := opand[len(opand)-1]
				if Level(op) < Level(expr[i]) {
					break
				} else {
					opand = opand[:len(opand)-1]
					if op == '(' {
						break
					}
					a, b := opnum[len(opnum)-2], opnum[len(opnum)-1]
					opnum[len(opnum)-2] = exec(a, b, op)
					opnum = opnum[:len(opnum)-1]
				}
			}
			if expr[i] != ')' {
				opand = append(opand, expr[i])
			}
		}
	}
	return opnum[0]
}

func Level(op byte) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	case '(', ')':
		return 0
	}
	return -1
}

func exec(a, b int, op byte) int {
	ret := 0
	if op == '+' {
		ret = a + b
	} else if op == '-' {
		ret = a - b
	} else if op == '*' {
		ret = a * b
	} else {
		ret = a / b
	}
	return ret
}

func main() {
	var expr string
	fmt.Scan(&expr)
	fmt.Println(Cacl(expr))
}
