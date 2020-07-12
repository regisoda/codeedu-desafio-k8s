package main

import "fmt"

func greeting(s string) string {
	return "<b>" + s + "</b>"
}

func main() {
	r := greeting("Code.education Rocks")
	fmt.Println(r)
}
