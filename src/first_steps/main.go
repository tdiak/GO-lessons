// Package name
package main

// imported modules
import (
	"fmt"
	"os"
)

// const value
const pi = 3.14



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


// Attribution example
func attributionExample(){
	var num1 int // standard attribution with default value 0
	var num2 = 10 // standard attribution
	var num3, bool1, string1 = 10, true, "some string" // multiple attribution
	i := 100 // short attribution

	fmt.Println("Attributions:")
	fmt.Println(num1) // 0
	fmt.Println(num2) // 10
	fmt.Println(num3) // 10
	fmt.Println(bool1) // true
	fmt.Println(string1) // some string
	fmt.Println(i) // 100
}


// Simple if - else example
func ifElse(first int){
	fmt.Println("If - Else example:")
	if first > 0{
		fmt.Println("number is highter than 0")
	} else if first == 0 {
		fmt.Println("number is equal to 0")
	} else {
		fmt.Println("number is lower than 0")
	}
}


// Simple loop example
func loopExample(items int){
	fmt.Println("Simple loop example:")
	for i := 0; i < items; i++ {
		fmt.Println(i)
	}
}


// Print command line args(using loop): to run main.go arg1 arg2 arg3
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
	attributionExample()
	ifElse(0)
	loopExample(7)

	printFileArgs()
}

// Excercise
// Create an application that converts celsius degrees into fahrenheit and kelvin
// Example input: go run main.go 100 12 40 123 30
// Output should be displayed on console

//
//
//
//
//
