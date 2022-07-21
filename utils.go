package utils

import "fmt"

func getOneNumber(say string) int {
	var a int
	fmt.Print(say)
	fmt.Scan(&a)
	return a
}
