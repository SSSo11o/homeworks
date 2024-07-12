package main
import "fmt"
func main (){
	var a float64
	var b float64
	var c float64
	fmt.Println(a, b, c)
	a = 2
	b = 3
	c = 4
	v := a * b * c
	s := 2 * (a * b + b * c + a * c) 
	fmt.Println(v, s)
}