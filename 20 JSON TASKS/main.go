package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//zadacha1
type Person struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}

// zadacha 5
type Product struct {
 Name string `json:"name"`
 Price float64 `json:"price,omitempty"`
 InStock bool `json:"in_stock,omitempty"`
}

func main() {
    // zadacha 1
    person := Person{
        Name:  "Aziz",
        Age:   21,
        Email: "ibronovaziz5@gmail.com",
    }
    a, b := json.Marshal(person)
    if b != nil {
        log.Fatalf("Ошибка при сериализации в JSON: %s", b)
    }
    fmt.Println(string(a))

    // zadacha2
    jsonstr := `{"name":"Aziz","age":21,"email":"ibronovaziz110@gmail.com"}`
    var person1 Person
    d := json.Unmarshal([]byte(jsonstr), &person1)
    if d != nil {
        log.Fatalf("Ошибка при десериализации JSON: %s", d)
    }
    fmt.Printf("Name: %s\nAge: %d\nEmail: %s\n", person1.Name, person1.Age, person1.Email)

    // zadacha 3
    date := map[string]int{
        "apples": 5,
        "oaranges": 10,
    }
    jsondat, j := json.Marshal(date)
    if j != nil {
        log.Fatalf("Ошибка при сериализации в JSON: %s", j)
        fmt.Println(string(jsondat))
    }

    // zadacha 5
    product := Product{
        Name: "Samsung",
    }
    s, r := json.Marshal(product)
    if r != nil {
        log.Fatalf("Ошибка при сериализации в JSON: %s", r)
    }
    fmt.Println(string(s))

    // zadacha 11
    person2 := Person {
        Name:  "Firuz",
        Age:   22,
        Email: "karomovf@gmail.com",
    }
    m, n := json.MarshalIndent(person2, "", "     ")
    if n != nil {
        log.Fatalf("Ошибка при сериализации в JSON: %s", n)
    }
    fmt.Println(string(m))

    // zadacha 14
    p := [] Person{
        {Name: "Umed", Age: 23, Email: "umedusupov01@gmail.com"},
        {Name: "Hasan", Age: 33, Email: "hasan0010@gmail.com"},
        {Name: "Murod", Age: 19, Email: "pulodovmurod1@gmail.com"},
    }
    k, l := json.Marshal(p)
    if l != nil {
        log.Fatalf("Ошибка при сериализации в JSON: %s", l)
    }
    fmt.Println(string(k))
}

    



