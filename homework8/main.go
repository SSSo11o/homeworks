package main
import "fmt"
func main (){
	// 1.Найти максимальный элемент в массиве.
	arr := [10]int {3,54,2,444,5,11,51,6,11,33}
	max := arr[0]
	for _, i := range arr[1 :] {
		if i > max {
			max = i
		}
	}
	fmt.Println(max)

	// 2.Найти минимальный элемент в массиве.
	arr1 := [10]int {3,54,2,444,5,11,-51,6,11,33}
	min := arr[0]
	for _, i := range arr1[1 :] {
		if i < min {
			min = i
		}
	}
	fmt.Println(min)

	// 3.Подсчитать количество положительных чисел в массиве.
	arr3 := [10]int {-3,54,2,-444,5,-11,-51,6,11,33}
	polozh := 0
	for _, i := range arr3[1 :] {
		if i > 0 {
			polozh++
		}
	}
	fmt.Println(polozh)

// 4.Найти сумму всех элементов массива.
arr4 := [10]int {44,54,2,444,5,13,-51,6,31,33}
sum := 0
for _, i := range arr4 {
	sum = sum + i
}
fmt.Println(sum)

// 5.Найти среднее значение всех элементов массива.
arr5 := [10]int {44,54,2,444,5,13,-51,6,31,33}
summ := 0
for _, i := range arr5 {
	summ = summ + i
}
sredZnach := float64(summ) / float64(len(arr5))
fmt.Println(sredZnach)

// 6.Удалить все вхождения заданного числа из массива.
arr6:= [10]int {3,54,2,444,5,11,-51,6,11,3}
num := 3
posle := []int{}
for _, m := range arr6{
	if m != num {
		posle = append(posle, m)
	}
}
fmt.Println(posle)

// 7.Умножить все элементы массива на заданное число.
arr7:= [10]int {3,54,2,444,5,11,-51,6,11,31}
um := 2
posle2 := []int{}
for _, n := range arr7{
	n = n * um
	posle2 = append(posle2, n)
}
fmt.Println(posle2)

// 8.Найти все индексы заданного числа в массиве.
    arr8 := [10]int{3, 4, 2, 4, 5, 11, -51, 4, 11, 4}
    num2 := 4
    indices := []int{}

    for i, m := range arr8 {
        if m == num2 {
            indices = append(indices, i)
        }
    }

    fmt.Println(indices)

// 9.Создать копию массива.
arr9 := [10]int{3, 90, 2, 66, 5, 11, 43, 4, 11, 1}
var copearr9 [10]int
for i, v := range arr9 {
	copearr9[i] = v
}
fmt.Println(copearr9)

// 10.Объединить два массива.
arr10 := []int{1,2,3,4,5}
ar10 := []int{7,8,9}
a10  := append(arr10, ar10...)
fmt.Println(a10)

// 11.Поменять местами максимальный и минимальный элементы массива.
arr11 := [7]int {4,1,5,65,-4,12,-44}
maxind, minind := 0, 0
for i := range arr11 {
	if arr11[i] > arr11[maxind] {
		maxind = i
	}
	if arr11[i] < arr11[minind] {
		minind = i
}
}
arr11[maxind], arr11[minind] = arr11[minind], arr11[maxind]
fmt.Println(arr11)

// 12.Проверить, является ли массив палиндромом.
arr12 := []int{1,2,2,1}
palindrome := true
for i := 0; i < len(arr12)/2; i++{
	if arr12[i] != arr12[len(arr12)-1-i]{
		palindrome = false
		break
	}
}
if palindrome {
	fmt.Println("Это палиндром")
}else {
	fmt.Println("не палиндром")
}

// 14.Перевернуть массив.
arr14 := []int{1, 2, 3, 4, 5}
for i, j := 0, len(arr14)-1; i < j; i, j = i+1, j-1 {
	arr14[i], arr14[j] = arr14[j], arr14[i]
}
fmt.Println(arr14)

}
