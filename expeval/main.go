package main

import (
	"fmt"
	"strconv"
)

func isDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

type Token struct {
	Ttype  int
	Lexeme string
}

const (
	literal    = iota //0
	plus       = iota //1
	minus      = iota //2
	leftparen  = iota //3
	rightparen = iota //4
)

func lexer(s string) []Token {
    var tokens []Token
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			tokens = append(tokens, Token{leftparen, "("})
			continue
		}
		if s[i] == ')' {
			tokens = append(tokens, Token{rightparen, ")"})
			continue
		}
		if s[i] == '+' {
			tokens = append(tokens, Token{plus, "+"})
			continue
		}
		if s[i] == '-' {
			tokens = append(tokens, Token{minus, "-"})
		}
		if s[i] == ' ' {
			continue
		}
		if isDigit(s[i]) {
			j := i
			for j < len(s) && isDigit(s[j]) {
				j++
			}
			tokens = append(tokens, Token{literal, s[i:j]})
			i = j-1
			continue
		}
	}
    return tokens
}
func parse(tokens []Token) int {
	i := 0

	var term func() int
	var primary func() int
	var unary func() int

	primary = func() int {
		if tokens[i].Ttype == literal {
			val, _ := strconv.ParseInt(tokens[i].Lexeme, 10, 32)
			i++
			return int(val)
		}
		if tokens[i].Ttype == leftparen {
			i++
			val := term()
			if (tokens[i].Ttype == rightparen) {
				i++
			} else {
				panic("no right paren")
			}
			return val
		}
		return 0
	}

	unary = func() int {
		if tokens[i].Ttype == minus {
			i++
			right := unary()
			return -right
		}
		return primary()
	}

	term = func() int {
		left := unary()
		for i < len(tokens) && (tokens[i].Ttype == plus || tokens[i].Ttype == minus) {
			op := tokens[i].Ttype
			i++
			right := unary()
			if op == plus {
				left += right
			} else {
				left -= right
			}
		}
		return left
	}

	return term()
}

func calculate(s string) int {
    tokens := lexer(s)
	return parse(tokens)
}

func main() {
	// // test := []string{"(3+19-(3-1-4+(9-4-(4-(1+(3)-2)-5)+8-(3-5)-1)-4)-5)"}
	// test := []string{"(19-(3-1-4+(9-4-(4-(1+(3)-2)-5)+8-(3-5)-1)-4)-5)"}
	// for _, s := range(test) {
	// 	fmt.Println(calculate(s))
	// }
	test2 := "1-(3+5-2+(3+19-(3-1-4+(9-4-(4-(1+(3)-2)-5)+8-(3-5)-1)-4)-5)-4+3-9)-4-(3+2-5)-10"
	fmt.Println(calculate(test2))
}
