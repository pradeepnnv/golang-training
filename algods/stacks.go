package algods

import "fmt"

/*
type Stack []string

func (s Stack) isEmpty() bool {
	return len(s) == 0
}

func (s *Stack) Pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	}

	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}
*/
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	fmt.Println("input-->", s)
	var stack Stack

	for _, c := range s {
		fmt.Printf("Workigng on --> %c\n", c)
		switch c {
		case '{':
			stack.Push(c)
			fmt.Printf("Stack after pushing %c is %v\n", c, stack)
		case '(':
			stack.Push(c)
			fmt.Printf("Stack after pushing %c is %v\n", c, stack)
		case '[':
			stack.Push(c)
			fmt.Printf("Stack after pushing %c is %v\n", c, stack)
		case '}':
			fmt.Println(stack)
			elem, ok := stack.Pop()
			fmt.Println(elem, ok)
			if !ok || elem != '{' {
				return false
			}
		case ')':
			fmt.Println(stack)
			elem, ok := stack.Pop()
			fmt.Println(elem, ok)
			if !ok || elem != '(' {
				return false
			}
		case ']':
			fmt.Println(stack)
			elem, ok := stack.Pop()
			fmt.Println(elem, ok)
			if !ok || elem != '[' {
				return false
			}
		}
	}
	return stack.IsEmpty()
}

type Stack []rune

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func (s *Stack) Push(r rune) {
	*s = append(*s, r)
	fmt.Printf("Stack after pushing: rune %c is %v\n", r, *s)
}

func (s *Stack) Pop() (rune, bool) {
	fmt.Println("Popping this stack", s)
	if (*s).IsEmpty() {
		return 0, false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}
