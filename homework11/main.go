package main

import (
	"fmt"
	"strings"
)

// 1. Конкатенация строк Напишите функцию, которая принимает две строки и возвращает их конкатенацию.
func concitensia(s1, s2 string) string {
	return s1 + s2
}

// 2.Длина строки Напишите функцию, которая принимает строку и возвращает её длину.
func dlina (d string)int{
	return len(d)

} 

// 4. Поиск подстроки Напишите функцию, которая находит индекс первого вхождения подстроки в строке. 	Верните -1, если подстрока не найдена.
func poisk(c, d string) int {
	return strings.Index(c, d)
} 

// 6.Удаление пробелов Напишите функцию, которая удаляет пробелы в начале и в конце строки.
func delate(f string)string{
	return strings.TrimSpace(f) // udalaet probeli
}
func main() {
// 1
	str1 := "Hello, "
	str2 := "world!"
	r := concitensia(str1, str2)
	fmt.Println( r)

// 2
str3 := "Hello Go"
lens := dlina(str3)
fmt.Println(lens)

//3 Проверка подстроки Напишите функцию, которая проверяет, содержит ли строка заданную подстроку.
s := "Hello, Aziz!"
fmt.Println(strings.Contains(s, "World")) 

// 4
str4 := "Humo, lab"
strr4 := "lab"
index := poisk(str4, strr4)
if index != -1 {
	fmt.Printf("Подстрока \"%s\" найдена в строке \"%s\" на индексе %d.\n", str4, strr4,  index)
} else {
	fmt.Printf("Подстрока \"%s\" не найдена в строке \"%s\".\n", str4, strr4)
}

// 6
str6 := "   Ibronov,   Aziz     "
delete := delate(str6)
fmt.Println("Исходная строка: '%s'\n", str6)
fmt.Println((delete))

// 7 

}


