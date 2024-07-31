//                          Metodi strukrura

package main

import "fmt"

type Rectangle struct {
	width, height float64
}
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main (){
	rect := Rectangle{width: 10, height: 5}
	area := rect.Area()
	fmt.Println(area)
}