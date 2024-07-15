package main

import (
	"fmt"
	"math"
	"strings"
)
func PrintGreeting(){
	fmt.Println("Hello, World!")
}
func DisplaySeparator(){
	fmt.Println("--------------------")
}

func NotifyStart(){
	fmt.Println("Program started")
}

func GetWelcomeMessage() string{
	return "Welcome!"
}

func GetPiValue()float64 {
	return math.Round(math.Pi*100) / 100
}

func IsServerActive()bool {
	return true
}
func PrintNumber (num int){
	fmt.Println(num)
}

func GreetUser(name string){
	fmt.Println("Hello", name)
}

func ToggleLight(Light bool){
	if Light {fmt.Println("Light on ")
		} else {
			fmt.Println("Light off ")
		}
}

func Add(a, b int)int{
	c := a + b
	return c
}

func Concat(st1, st2 string) string {
	kon := st1 + st2
	return kon
}

func IsEven(num1 int) bool {
	if num1 % 2 == 0 {
		return true
	} else {
		return false
	}
}

//Создайте функцию Calculate, которая принимает два целых числа и функцию для их обработки. Примените переданную функцию к этим числам и верните результат.

func Calculate(a, b int, culate func (int, int) int) int {
	return culate (a, b)
}
func plas (a, b int) int {
	return a + b
}
func minus (a, b int) int {
	return a - b
}

// •	Создайте функцию Execute, которая принимает булевое значение и функцию, которая принимает булевое значение и ничего не возвращает. Выполните переданную функцию с переданным булевым значением.
func Execute (znach bool, action func (bool))  {
	action (znach)
}
func prent (a bool) {
	if a {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

// •	Создайте функцию Apply, которая принимает целое число и функцию, которая принимает целое число и возвращает целое число. Примените переданную функцию к переданному числу и верните результат.
func Apply (a float64, action func(int)int) int {
	inVal := int(math.Floor(a)) 
	return action(inVal)
}
func plus (num int) int {
    return num + 1
}

// •	Создайте функцию Multiplier, которая принимает целое число factor и возвращает функцию, умножающую переданное ей число на factor.
func Multiplier (factor int) func(int) int {
	return func(a int) int {
        return a * factor
	}
}

// •	Создайте функцию StringRepeater, которая принимает целое число n и возвращает функцию, повторяющую переданную ей строку n раз.
func StringRepeater(n int) func(string) string {
    return func(s string) string {
        return strings.Repeat(s, n)
    }
}

// •	Создайте функцию ConditionalPrinter, которая принимает булево значение condition и возвращает функцию, печатающую строку, если condition истинно.
func ConditionalPrinter (condition bool) func (string) {
	return func(s string) {
		if condition {
			fmt.Println(s)
		}
	}
}

func main(){
	PrintGreeting()
	DisplaySeparator()
	NotifyStart()

	WelcomeMessage := GetWelcomeMessage()
	fmt.Println(WelcomeMessage)

	PiValue := GetPiValue()
	fmt.Println(PiValue)

	ServerActive := IsServerActive()
	fmt.Println(ServerActive)

	PrintNumber(14)
	GreetUser("Aziz")

	ToggleLight(true)
	ToggleLight(false)

	fmt.Println(Add(4, 6))

	fmt.Println(Concat("Ibronov", "__a"))

	fmt.Println(IsEven(4))
	fmt.Println(IsEven(7))

	// •	Создайте переменную adder, которая является функцией, принимающей два целых числа и возвращающей их сумму.
	adder := func (a, b int) int {
		return a + b
	}
	adde := adder(3, 5)
	fmt.Println(adde)

	// •	Создайте переменную concatenator, которая является функцией, принимающей две строки и возвращающей их конкатенацию.
	concatenator := func (st1, st2 string) string {
		return st1 + st2
	}
	con := concatenator("Go", "Land")
	fmt.Println(con)

	// •	Создайте переменную isPositive, которая является функцией, принимающей целое число и возвращающей true, если число положительное.
	isPositive := func (num int) bool {
		if num > 0 { 
			return true
		} else {
			return false
		}
	}
	Positive := isPositive(4)
	fmt.Println(Positive)

// •	Создайте функцию Calculate, которая принимает два целых числа и функцию для их обработки. Примените переданную функцию к этим числам и верните результат.
	resalt := Calculate(23, 12, plas)
	fmt.Println(resalt)
	resalt = Calculate(25, 5, minus)
	fmt.Println(resalt)

// •	Создайте функцию Execute, которая принимает булевое значение и функцию, которая принимает булевое значение и ничего не возвращает. Выполните переданную функцию с переданным булевым значением.
Execute(true, prent)
Execute(false, prent)
result := Apply(5.8, plus)
    fmt.Println( result)

	// Создайте функцию Multiplier, которая принимает целое число factor и возвращает функцию, умножающую переданное ей число на factor.
	b := Multiplier(3)
	resalt = b(5)
	fmt.Println(resalt)

	// •	Создайте функцию StringRepeater, которая принимает целое число n и возвращает функцию, повторяющую переданную ей строку n раз.

	// // •	Создайте функцию ConditionalPrinter, которая принимает булево значение condition и возвращает функцию, печатающую строку, если condition истинно.
	printTrue := ConditionalPrinter(true)
	printTrue("rabotaet znachenie true")
}