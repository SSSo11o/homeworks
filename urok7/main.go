//                                Proizvodnie tipi
package main
import "fmt"
type Age int   // Age budet int  toest u Age budet svoestva int
func main (){
var myAge Age = 25
fmt.Println(myAge)

fmt.Println(a(14))

fmt.Println(chislo(23))

fmt.Println(otsenka(101))

obratnie(9)

temperatur(4)

fmt.Println(sena(78))

user1 := User{Name: "Aziz", Age: 21} 
fmt.Println(chelovek(user1))

/* struktura 
p := Person(name: "aziz", Age: 30)
ukazatel p := &person(name: "aziz", Age: 30)
 */
}

// psevdonimi tupe Myint = int -- eto tozhe samoe kak vishe(proizvodnie tip) no v etom cvoestva ne izmenieetsa 

/*                           Struktura
type Person struct {
Name string
Age int
}  */

/* Vlojenie 
type contact struct{
    email string
    phone string
}
 
type person struct{
    name string
    age int
    contactInfo contact
} 
func main() {
     
    var tom = person {
        name: "Tom", 
        age: 24,
        contactInfo: contact{
 email: "tom@gmail.com",
            phone: "+1234567899",
        },
    }
    tom.contactInfo.email = "supertom@gmail.com"
     
    fmt.Println(tom.contactInfo.email)
    fmt.Println(tom.contactInfo.phone)  */

	//               Хранения ссылки на структуру того же типа ---- ?

	/*                            Zadachi  
	 Определение возраста совершеннолетия
Определите тип Age на основе int. Напишите функцию, которая принимает возраст и возвращает сообщение о том,
 является ли человек совершеннолетним (возраст 18 лет и старше) или нет */

type age int
 func a(age1 age) string {
	if age1 >= 18 { 
		return "совершеннолетные"
		
	} else {
		return "несовершеннолетные"
	}
}

/* // Проверка на четность
 Определите тип Number на основе int. Напишите функцию, которая принимает число и возвращает сообщение о том, является ли оно четным или нечетным */
type Number int
func chislo(number1 Number) string {
	if number1 % 2 == 0 {
		return "четным"
	} else {
		return "нечетным"
	}
}

/* Проверка допустимого диапазона
// Определите тип Score на основе int. Напишите функцию, которая принимает оценку и возвращает сообщение, находится ли она в допустимом диапазоне от 0 до 100 */
type Score int 
func otsenka (scor1 Score) string {
	if scor1 >= 0 && scor1 <= 100 {
		return "находится в диапазоне"
	} else {
		return "не находится в диапазоне"
	}
}

/* Обратный отсчет
Создайте псевдоним для типа int под названием Count. Напишите функцию, которая принимает Count и выводит обратный отсчет до нуля */
 type Count = int
 func obratnie(count1 Count){
	for i := count1; i > 0; i--{
		fmt.Println(i)
	}
 }

 /* Проверка температуры
// Создайте псевдоним для типа float64 под названием Temperature. Напишите функцию, которая принимает Temperature и выводит сообщение о состоянии (ниже нуля, выше нуля или ноль). */
type Temperature = float64
func temperatur(temp Temperature) {
	if temp < 0 {
		fmt.Println("ниже нуля")
	} else if temp > 0 {
		fmt.Println("выше нуля")
	} else if temp == 0 {
		fmt.Println("ноль")
	}
}

/* Применение скидки
// Создайте псевдоним для типа float64 под названием Price. Напишите функцию, которая принимает Price и возвращает новую цену с 10% скидкой. */
type Price = float64
func sena (price1 Price) Price{
	price1 = price1 - price1 * 0.1 
	return price1
}

/* Информация о пользователе
Создайте структуру User с полями Name (строка) и Age (целое число). Напишите функцию, которая принимает User и выводит информацию о нем */
type User struct {
	Name string
	Age int
}
func chelovek (users User) string {
	str := fmt.Sprintf("Имя: %s, возраст: %d", users.Name, users.Age)
	return str
}