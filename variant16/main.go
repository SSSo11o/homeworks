// Вариант 16
package main
import "fmt"

//Напишите функцию, которая принимает срез целых чисел и возвращает новый срез, содержащий только те числа, которые являются квадратами чисел.
func kvadrat(n int) bool {
	if n < 0 {
		return false
	}
	for i := 0; i*i <= n; i++ {
		if i*i == n {
			return true
		}
	}
	return false
}
func filterPerfectSquares(nums []int) []int {
	var a []int
	for _, num := range nums {
		if kvadrat(num) {
			a = append(a, num)
		}
	}
	return a
}

func main (){
	nums := []int{1, 2, 3, 4, 5, 9, 16, 20, 25, 30}
	a := filterPerfectSquares(nums)
	fmt.Println(a)

	// Реализуйте функцию, которая принимает срез строк и возвращает новый срез, содержащий только те строки, которые являются палиндромами.
	arr := []string{"hi", "Aziz", "GO", "Aziz", "hi"}
	palindrome := true
	for i := 0; i < len(arr)/2; i++{
		if arr[i] != arr[len(arr)-1-i]{
			palindrome = false
			break
		}
	}
	if palindrome {
		fmt.Println("Это палиндром")
	}else {
		fmt.Println("не палиндром")
	}
	
}
