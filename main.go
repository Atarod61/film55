package main

import "fmt"

 
 func Print[T any](value T)T{
 func Print[T any](value T) T {
 	fmt.Println(value)
 	return value
 }
 
 func main() {
 	//peintstring("hello world")
 	//printinteger(12)
 }
 //func peintstring(s string) string {
 	fmt.Println(s)
 	return s
 
 //}
 	Print(12)
 	Print("hi my name is alireza")
 	panic(12.5)
 
 //func printinteger(i int) int {
 	//fmt.Println(i + 100)
 	//return i
 //}
 }
