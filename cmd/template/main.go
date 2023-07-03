package main

import "github.com/tanveerprottoy/go-gin-template/internal/app/template"

func main() {
	a := template.NewApp()
	a.Run()
}

// Multiply just to check unit test
func Multiply() int {
	return 25 * 4
}
