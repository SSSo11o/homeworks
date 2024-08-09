package main

import "fmt"

func main() {
	book1 := Book{title: "Tojikon", aurhor: "Bobojon Gafurov"}
	book2 := Book{title: "Shohnoma", aurhor: "Firdavsi"}
	var library Library
	library.AddBook(book1)
	library.AddBook(book2)
	for _, book := range library.books {
		fmt.Println(book.GetDetails())
	}

	r := Circle{radius: 5}
	fmt.Printf("Area: %.2f\n", r.Area())
	fmt.Printf("Circumference: %.2f\n", r.Circumference())

	temp := Temperature{celsius: 32}
	fmt.Printf("Celsius: %.2f\n", temp.celsius)
	fmt.Printf("Fahrenheit: %.2f\n", temp.ToFahrenheit())
	fmt.Printf("Kelvin: %.2f\n", temp.ToKelvin())
}

/* 1.Книга и Автор
Описание: Реализуйте структуру Book с полями title и author, а также метод GetDetails, который возвращает строку с деталями книги. Реализуйте структуру Library с массивом книг и метод AddBook, добавляющий книгу в библиотеку.
Методы:
GetDetails() string для структуры Book
AddBook(book Book) для структуры Library */

type Book struct {
	title, aurhor string
}

func (b Book) GetDetails() string {
	return fmt.Sprintf("Title: %s, Author: %s", b.title, b.aurhor)
}

type Library struct {
	books []Book
}

func (l *Library) AddBook(book Book) {
	l.books = append(l.books, book)
}

/* 3.Круг и Площадь
Описание: Реализуйте структуру Circle с полем radius. Реализуйте методы Area и Circumference для вычисления площади и периметра круга.
Методы:
Area() float64
Circumference() float64*/
type Circle struct {
	radius float64
}

const Pi = 3.141592653589793

func (c Circle) Area() float64 {
	return Pi * c.radius * c.radius
}

func (c Circle) Circumference() float64 {
	return 2 * Pi * c.radius
}

/* 7.Работа с температурой
Описание: Реализуйте структуру Temperature с полем celsius. Реализуйте методы для преобразования температуры в Фаренгейт и Кельвин.
Методы:
ToFahrenheit() float64
ToKelvin() float64 */
type Temperature struct {
	celsius float64
}

func (t Temperature) ToFahrenheit() float64 {
	return t.celsius*9/5 + 32
}
func (t Temperature) ToKelvin() float64 {
	return t.celsius + 273.15
}
