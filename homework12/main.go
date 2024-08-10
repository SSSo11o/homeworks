package main
import "fmt"
func main (){
	// 1.Создание и вывод map Описание: Создайте map для хранения строк и их длин, добавьте несколько элементов и выведите содержимое.
vivod := make(map[string]int)
vivod["Ibronov"] = len("Ibronov")
vivod["Aziz"] = len("vivod")
for ctrok, dlina := range vivod {
	fmt.Printf("Строка: %s, Длина: %d\n", ctrok, dlina)
}

// 2.Проверка наличия ключа Описание: Создайте map с несколькими элементами и напишите функцию, которая проверяет наличие заданного ключа.
element := map[string]int{
	"Go": 2,
	"HTML": 4,
	"CSS": 3,
}
key := "Go"
if _, a := element[key]; a {
	fmt.Println("Ключ", key, "присутствует в map.")
} else {
	fmt.Println("Ключ", key, "отсутствует в map.")
}

// 3.Обновление значения по ключу Описание: Напишите функцию для обновления значения в map по заданному ключу.
element2 := map [string]int{
	"Hello": 5,
	"Go": 2,
	"Humo": 4,
}
b := "Humo" 
element2[b] = 10
fmt.Println(element2)

// 4
element3 := map[string]int{
	"dog": 3,
	"cats": 4,
	"elephant": 8,
}
c := "dog"
	deleteElement(element3, c)
	c = "Unknown"
	deleteElement(element3, c)
	fmt.Println(element3)

	// 7
	pustomap := map[string]int {}
	element4 := map[string]int {
		"Hello": 5,
        "World": 5,
	}
	fmt.Println ("pustomap пустой?",pusto(pustomap))
	fmt.Println("element4 пустой?", pusto(element4))
}

//4.Удаление элемента из map Описание: Создайте функцию, которая удаляет элемент из map по заданному ключу.
func deleteElement(m map[string]int, c string) {
    delete(m, c)
    fmt.Printf("Ключ '%s' удалён из map.\n", c)
}

// 7.Проверка пустого map Описание: Напишите функцию, которая проверяет, пустой ли map.
func pusto(e map[string]int) bool{
	return len(e) == 0
}