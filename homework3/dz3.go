package main
import "fmt"

// 1.	Создайте функцию PrintGreeting, которая печатает "Доброе утро!", если dayType равен "утро"; "Добрый день!", если dayType равен "день"; и "Добрый вечер!", если dayType равен "вечер". Переменную  dayType вводить с консоли внутри функции.
func PrintGreeting ()  {
	 dayType := "утро"
switch dayType {
case "утро":
	fmt.Println("Доброе утро!")
case "день":
	fmt.Println("Добрый день!")
case "вечер":
	fmt.Println("Добрый вечер!")
}
}

// 2.	Создайте функцию PrintWeather, которая печатает "Солнечно", если weatherType равен "солнечно"; "Облачно", если weatherType равен "облачно"; и "Дождливо", если weatherType равен "дождливо". Переменную  weatherType вводить с консоли внутри функции.
func PrintWeather () { 
 weatherType := "солнечно" 
switch weatherType {
case "солнечно" :
	fmt.Println("Солнечно")
case "облачно"  :
	fmt.Println("облачно")
case "дождливо" :
	fmt.Println("дождливо")
}
}

// 3.	Создайте функцию PrintTrafficLight, которая печатает "Стоп", если lightColor равен "красный"; "Внимание", если lightColor равен "желтый"; и "Идите", если lightColor равен "зеленый". Переменную lightColor вводить с консоли внутри функции.
func PrintTrafficLight (){
	lightColor := "красный"
	switch lightColor {
	case "красный" :
		fmt.Println("Стоп")
	case "желтый" : 
	fmt.Println("Внимание")
	case "зеленый" :
		fmt.Println("Идите")
	}
}

// Создайте функцию GetGrade, которая возвращает оценку "A", "B" или "C" в зависимости от значения переменной score. Переменную scope вводить с консоли внутри функции.
func GetGrade (){
	score := 4 
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 75 {
		fmt.Println("B")
	} else if score < 75 {
		fmt.Println("C")
	}
}

// 5.Создайте функцию GetDiscount, которая возвращает скидку "10%" - при amount > 1000, "5%" - при 500 < amount <=100  или "0%" - при amount <= 500 в зависимости от суммы покупки amount. Переменную amount вводить с консоли внутри функции.
func GetDiscount (){
	amount :=  700
	if  amount > 1000 {
		a := amount * 10 / 100 
		fmt.Println("у вас скидка :", a)
	} else if 500 < amount && amount <= 1000 {
		a := amount * 5 / 100
		fmt.Println("у вас скидка :", a)
	} else if amount <= 500 {
		fmt.Println("покупайте больше")
	} 
}

// Создайте функцию GetTemperatureDescription, которая возвращает "Холодно" (меньше 10) , "Тепло" (с 10 до 25)  или "Жарко" в зависимости от значения переменной temperature. Переменную temperature вводить с консоли внутри функции.
func GetTemperatureDescription () {
	temperature := 5
	if temperature < 10 {
		fmt.Println("Холодно")
	} else if temperature >= 10 && temperature < 25 {
		fmt.Println("Тепло")
	} else if temperature > 25 {
		fmt.Println("Жарко")
	}
}

// Создайте функцию CheckNumber, которая принимает целое число и печатает "Положительное", если число больше нуля; "Отрицательное", если меньше нуля; и "Ноль", если равно нулю.
func CheckNumber () {
	var a int 
	a = 5
	if a > 0 {
		fmt.Println("Положительное")
	} else if a < 0 {
		fmt.Println("Отрицательное")
	} else if a == 0 {
		fmt.Println("Ноль")
	}
}

// 8.Создайте функцию CheckAge, которая принимает возраст и печатает "Совершеннолетний", если возраст 18 и старше; "Несовершеннолетний", если младше 18.
func CheckAge () {
	a := 18
	if a >= 18 {
		fmt.Println("Совершеннолетний")
	} else if a < 18 {
		fmt.Println("Несовершеннолетний")
	}
}

// 9.Создайте функцию CheckPassword, которая принимает строку пароля и печатает "Пароль верный", если пароль равен "1234"; и "Пароль неверный" в противном случае.
func CheckPassword (){
	a := 1234
	if a == 1234 {
		fmt.Println("Пароль верный")
	} else {
		fmt.Println("Пароль неверный")
	}
}

// 10.Создайте функцию Add, которая принимает два целых числа и возвращает сумму их модулей. Используйте условные операторы для проверки значений чисел.
func Add (a, b int) int {
	if a < 0 {
		a = -a 
	} 
	if b < 0{
		b = -b
	} 
	return a + b
}

// 11. Создайте функцию CompareStrings, которая принимает две строки и возвращает "равны", если строки одинаковы, и "не равны" в противном случае. Используйте условные операторы для сравнения строк.
func CompareStrings (str1, str2 string) string {
	if str1 == str2 {
	return "равны"
} else {
	return "равны"
}
}

// 12 Создайте функцию Max, которая принимает два целых числа и возвращает большее из них. Используйте условные операторы для сравнения чисел.
func Max (c, d int) int{
if c > d {
	return c
} else if d > c {
	return d
} else {
	return 0
}
}

// 13.Создайте переменную operation, которая является функцией, принимающей два целых числа и возвращающей их разность. Используйте условные операторы внутри функции для проверки значений чисел.
// operation := func (e, f int) int {
// 	g := e - f
// 	if g < 0 {
// 		g = -g
// 	}
// 	return g
// }
func main() {
	PrintGreeting()
	PrintWeather()
	PrintTrafficLight()
	GetGrade()
	GetDiscount()
	GetTemperatureDescription()
	CheckNumber()
	CheckAge()
	CheckPassword()
	println(Add(-5,-7))
	println(CompareStrings("Aziz", "Ibronov"))
	println(Max(4,6))
	// println(operation(3, 5))
}

