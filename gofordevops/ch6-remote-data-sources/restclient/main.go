package main

import "fmt"

type getReq struct {
	Author string `json:"author`
}

type getQuote struct {
	Quote string `json:"quote"`
	Error *Error `json:"error"`
}

type Error struct {
	Code int
	Msg  string
}

func (e *Error) Error() string {
	return fmt.Sprintf("(code %v): %s", e.Code, e.Msg)
}

func main() {

}
