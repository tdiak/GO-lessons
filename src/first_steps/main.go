// Package name
package main

// imported modules
import (
	"fmt"
	"os"
)


// Simple calculator functions
func add(x int, y int) int{
	return x + y
}

func diff(x int, y int) int{
	return x - y
}

func multiplication(x int, y int) int{
	return x * y
}

func division(x int, y int) int{
	return x / y
}


// Print command line args(using loop): fo run main.go arg1 arg2 arg3
func printFileArgs(){
	var s, sep string
	for i := 1; i < len(os.Args); i++{
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}


// main function
func main(){
	fmt.Println("Hi, this is your first lesson")
	fmt.Println("Calculator example(8 and 4):")
	fmt.Println(add(8, 4))
	fmt.Println(diff(8, 4))
	fmt.Println(multiplication(8, 4))
	fmt.Println(division(8, 4))



	printFileArgs()
}