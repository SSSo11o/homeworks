package main
import "fmt"
func sum(a, b int) int {
	c := a + b
	return c
}

func printhello(name string) {
	fmt.Println("Hello", name)
}

func printhello2(){
	fmt.Println("Hello")
}

func inc(a int ){
	a++
	fmt.Println(a)
}

func sum2 (a int) func() int {
	b := 6
	return func() int {
		return 	a + b
	}
}
func main (){
	fmt.Println(sum(7, 6))
	fmt.Println(sum(4, 1))
	c := 8
	d := float64(c)
	fmt.Println(d)
	d = d + 0.04
	fmt.Println(d)


	printhello("Aziz")
	printhello2()
	a := 7
	inc(a)
	fmt.Println(a)

	f := sum2(9)
	i := f()
	fmt.Println(i)
}