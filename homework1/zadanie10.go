package main
import "fmt"
func main (){
	var a, b float64
	fmt.Println(a, b)
	a = 4
	b = 3
	s := a*a + b*b
    r := a*a - b*b
    p := a * a * b * b
    d := (a * a) / (b * b)
	fmt.Println(s, r, p, d)

}