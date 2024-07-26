/*                   Masiv
Массивы представляют последовательность элементов определенного типа
var numbers = [...]int{1,2,3,4,5}   // длина массива 5
numbers2 := [...]int{1,2,3}         // длина массива 3
fmt.Println(numbers)                // [1 2 3 4 5]
fmt.Println(numbers2)           // [1 2 3]           */

package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5}
	slice3 := append(slice2, slice1...)
	fmt.Println(slice3)
	fmt.Println(len(slice3), cap(slice3))

	//                               Zadachi
	// Создайте массив из 7 целых чисел и инициализируйте его значениями от 1 до 7
	var numbers = [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(numbers)

	// Создайте массив из 5 строк и инициализируйте его значениями "a", "b", "c", "d", "e".
	var number1 = [...]string{"a", "b", "c", "d", "e"}
	fmt.Println(number1)

	// Создайте массив из 4 логических значений и инициализируйте его значениями true, false, true, false.
	var bol = [4]bool{true, false, true, false}
	fmt.Println(bol)

	// Создайте массив из 10 целых чисел без инициализации и выведите его на экран
	var a [10]int
	fmt.Println(a)

	// Создайте массив из 4 логических значений и выведите значения по индексам 1 и 3
	var b [4]bool
	b[0] = true
	b[1] = false
	b[2] = true
	b[3] = false
	fmt.Println(b[1])
	fmt.Println(b[3])

	//               ZADACHA

	d := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11}
	idx := 0
	for i, v := range d {
		if d[0] < v && v < d[9] {
			idx = i
		}
	}
	fmt.Println(idx)
}
