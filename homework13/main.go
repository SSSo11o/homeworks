package main

import (
	"fmt"
	"strings"
)
func main () {
	//1
stroka := MyString{text: "Hello Teacher"}
fmt.Printf("Текст: \"%s\"\n", stroka.text)
fmt.Printf("Количество слов: %d\n",  stroka.WordCount())
fmt.Printf("Длина строки: %d\n", stroka.Length())

//2
f := MyFormatter{texts: "Aziz Ibronov"}
fmt.Printf("Строка: %s\n", f.texts)
	fmt.Printf("Upper: %s\n", f.ToUpper())
	fmt.Printf("Lower: %s\n", f.ToLower())

	//3
	d := 10
	point := IntPointer{b: &d}
	fmt.Printf("значение: %d\n", *point.b) 
	point.Increment()
	fmt.Printf("После Increment: %d\n", *point.b)
	point.Decrement()
	fmt.Printf("После Decrement: %d\n", *point.b)

	// 4
	text := TextCleaner{text: "      Go   lang    "}
	fmt.Printf("Строка: \"%s\"\n", text.text)
	fmt.Printf("После TrimSpaces: \"%s\"\n", text.TrimSpaces())
	fmt.Printf("После RemoveSpaces: \"%s\"\n", text.RemoveSpaces())
}
/* 1.  Длина строки и Количество слов
Описание: Реализуйте интерфейс StringProcessor, который будет содержать методы Length() и WordCount(). Реализуйте структуру MyString, которая будет работать с текстом и реализует этот интерфейс.
Методы:
Length() int для получения длины строки.
WordCount() int для подсчета количества слов */
type StringProcessor interface{
	lenght() int
	WordCount () int
}
type MyString struct{
	text string
}
func (m MyString) Length() int {
	return len(m.text)
}
func (m MyString) WordCount() int {
	return len(strings.Fields(m.text))
}

/* 2 Форматирование строки
Описание: Реализуйте интерфейс Formatter с методами ToUpper() и ToLower(). Реализуйте структуру MyFormatter, которая будет работать со строкой и реализует этот интерфейс.
Методы:
ToUpper() string для преобразования строки в верхний регистр.
ToLower() string для преобразования строки в нижний регистр. */
type Formatter interface {
	ToUpper() string
	ToLower() string
}
type MyFormatter struct {
	texts string
} 
func (mm MyFormatter) ToUpper() string {
	return strings.ToUpper(mm.texts)
}
func (mm MyFormatter) ToLower() string {
	return strings.ToLower(mm.texts)
}

/*  3. Работа с указателями на числа
Описание: Реализуйте интерфейс PointerOperations с методами Increment() и Decrement(). Реализуйте структуру IntPointer с указателем на число, которая реализует этот интерфейс.
Методы:
Increment() для увеличения значения числа на 1.
Decrement() для уменьшения значения числа на 1. */
type PointerOperations interface {
	Increment() 
	Decrement() 
}
type IntPointer struct {
	b *int
}
func (i *IntPointer) Increment() {
	*i.b++
}
func(i *IntPointer) Decrement() {
	*i.b--
}

/* 4. Удаление пробелов в строке
Описание: Реализуйте интерфейс StringCleaner с методами TrimSpaces() и RemoveSpaces(). Реализуйте структуру TextCleaner, которая будет работать со строками и реализует этот интерфейс.
Методы:
TrimSpaces() string для удаления пробелов с начала и конца строки.
RemoveSpaces() string для удаления всех пробелов из строки.*/
type StringCleaner interface {
	TrimSpaces() string
	RemoveSpaces() string
}
type TextCleaner struct{
	text string
}
func (t TextCleaner)TrimSpaces()string{
	return strings.TrimSpace(t.text)
}  
func (t TextCleaner) RemoveSpaces() string {
	return strings.ReplaceAll(t.text, " ", "")

}