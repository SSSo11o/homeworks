package main
import "fmt"
func main (){
	// for i := 1; i <=10; i ++ {
	// 	fmt.Println(i)
	// }
	l := 1
	for i := 1; i <=5; i++{
		l = i * i
		fmt.Println(l)
	}
	m := 1
	for i := 1; i <= 10; i++ {
		m = i *3 
		fmt.Println(m)
	}

	// последовательности Фибоначчи      
    a, b := 0, 1  
    fmt.Println(a)
    fmt.Println(b)
    for i := 2; i < 10; i++ {
        c := a + b
        fmt.Println(c)
        a, b = b, c 
    }

	fmt.Println(evk(48,18))

	// 6.Напишите программу, которая выводит числа от 1 до 100, заменяя числа, кратные 3, на "Fizz", числа, кратные 5, на "Buzz", а числа, кратные 3 и 5, на "FizzBuzz".
	for i := 1; i <= 100; i++ { 
		switch {
        case i%3 == 0 && i%5 == 0:
            fmt.Println("FizzBuzz")
        case i%3 == 0:
            fmt.Println("Fizz")
        case i%5 == 0:
            fmt.Println("Buzz")
		default:
            fmt.Println(i)
	}
}

fmt.Println(pros(7))
resalt := ssum(42)
fmt.Println(resalt)

// 10.Напишите программу, которая запрашивает у пользователя ввод положительного числа и повторяет запрос, пока не будет введено положительное число.
var number int
for { fmt.Scan(&number)
	if number < 0 {
		fmt.Println("введите положительное число")
		continue
	}else {
		fmt.Println("вы ввели положительное число")
		break
	}
}

// 11.Напишите программу, которая вычисляет произведение всех чисел от 1 до введённого числа n, но прекращает выполнение, если результат превышает 1000.
var z int
z = 1
n := 10
for i :=1; i < n; i++{
	z = z * i
	if i > 1000 {
		break
	} else {
		fmt.Println(z)
		continue
	}
}

//12.  Напишите программу, которая запрашивает у пользователя ввод числа и проверяет, является ли оно палиндромом.
var k int 
k = 5
var p int 
p = 4
for {
	if k == p {
		fmt.Println("числы ровны")
		break
	} else {
		fmt.Println("повториты")
		continue
	}
}

// 14.Напишите программу, которая выводит таблицу умножения от 1 до n, где n - введенное пользователем число.
 var h int
 var j int
h = 7
j = 8
for l := 1; l < h; l++ {
 	for k := 1; k < j; k++ {
 		fmt.Printf("%d * %d = %d\n", l, k, l * k)
 	} 
 	fmt. Println("----------------")
 }

// Напишите программу, которая бесконечно запрашивает у пользователя ввод двух чисел и выводит их сумму
 var num1, num2 int
	for {
		fmt.Print("Введите первое число: ")
		fmt.Scan(&num1)
		fmt.Print("Введите второе число: ")
		fmt.Scan(&num2)
		sum := num1 + num2
		fmt.Printf("Сумма чисел %d и %d равна %d\n", num1, num2, sum)
	}
}


	//   алгоритм Евклида
	func evk(d, c int) int {
		for c != 0 {
			d, c = c, d % c
		}
		return d
	}
//             Напишите программу, которая проверяет, является ли число простым.
	func pros(n int) bool {
		if n <= 1 {
			return false
		}
		if n == 2 {
			return true
		}
		for i := 2; i < n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	// 9.Напишите программу, которая находит сумму цифр числа.
	func ssum(n int) int {
		var sum int
		sum = 0
		for i := 1; i <= n; i++{
			sum = sum + i
		}
		return sum
	}

	/* 15.Напишите программу, которая выводит треугольник Паскаля высотой n, где n - введенное пользователем число.
	func factorial(n int) int {
		resalt := 1
		for i := 1; i <= n; i++{
			resalt *= 1
		}
		return resalt
	}

	func binomialcoe(n, k int) int {
		return factorial()
	} */