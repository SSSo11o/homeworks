package main
import "fmt"
func main (){
fmt.Println(temp(4))
fmt.Println(vozrast(16))
fmt.Println(skorost(134))
fmt.Println(sscore(-5))
fmt.Println(opeeration(5))
fmt.Println(operationn(21))
fmt.Println(chekk(11))
fmt.Println(cheslo(1))
fmt.Println(plas(12, 5))
fmt.Println(minas(54, 9))
fmt.Println(umnozh(3, 6))
otshet(7)
fmt.Println(level(48))
fmt.Println(ves(99))
fmt.Println(osenka(54))
fmt.Println(zdorovie(157,78))

product2 := Product{Name: "potato", Price: 120}
fmt.Println(produc(product2))

person1 := Person{FirstName: "Aziz", LastName: "Ibronov", Age: 21}
fmt.Println(people(person1))

product5 := Product{Name: "apple", Price: 45}
product6 := Product{Name: "banana", Price: 55}
fmt.Println(prodyct(product5, product6))

product7 := Product{Name: "Rolex", Price: 1500}
fmt.Println("Before:", product7)
proddduct(&product7, 1850)
fmt.Println("after", product7)

person5 := Person{FirstName: "Firuz", Age: 21}
fmt.Println("Before:", person5)
personn(&person5,)
fmt.Println("after:", person5)

product11 := Product{Name: "Galaxy", Price: 4000}
fmt.Println("before:", product11)
ppproduct(&product11, "Redmi", 4500)
fmt.Println("after:", product11)

kolichestvo :=Item {Name: "apple", Quantity: 12}
fmt.Println("before:", kolichestvo)
items(&kolichestvo, 3)
fmt.Println("after:", kolichestvo)

persen := Persen{
	Name: "Aziz",
	Address: Address{
		Street: "Mekhkalon",
		City: "Dushanbe",
	},
}
adres(persen)

compane := Company {
	Name: "Humo",
	Location: Address{
		Street: "Operka",
		City: "Dushanbe",
	},
}
company(compane)

corse := Course{
	Title: "Learn GOlang",
	Instructor: Persen{
		Name: "Halimjon",
		Address: Address{
			Street: "Operka",
			City: "Dushanbe",
		},
	},
}
cours(corse)

struc := Event{
	Title: "Hello Go",
	Location: Address{
		Street: "Ashan",
		City: "New-york",
	},
}
even(struc)

projec := Project {
	Name: "Samsung",
	Manager: Persen{
		Name: "Azam",
		Address: Address{
			Street: "Zarafshon",
			City: "Dushanbe",
		},
	},
}
project(projec)
}


/* 1.Проверка температуры
Определите тип Temperature на основе float64. Напишите функцию, которая принимает температуру и возвращает сообщение о том, является ли она ниже, выше или равной нулю. */
type Temperature = float64
func temp(temperatur1 Temperature)string {
	if temperatur1 < 0 {
		return "ниже нуля"
	} else if temperatur1 > 0 {
		return "выше нуля"
	} else {
		return "ноль"
	}
}

/* 2.Проверка возраста
Определите тип Age на основе int. Напишите функцию, которая принимает возраст и возвращает сообщение о том, является ли человек подростком (возраст от 13 до 19 лет включительно) или нет. */
type Age int
func vozrast(age1 Age) string {
	if age1 >= 13 && age1 <= 19 {
		return "человек является подростком"
	} else {
		return "нет"
	}
}

/* 3.Проверка скорости
Определите тип Speed на основе float64. Напишите функцию, которая принимает скорость и возвращает сообщение о том, является ли она допустимой (от 0 до 120 км/ч включительно) или нет.*/
type Speed float64
func skorost (sped1 Speed) string {
	if sped1 >= 0 && sped1 <= 120 {
		return "скорость допустима"
	} else {
		return "нет"
	}
}

/* 4.Проверка счета
Определите тип Score на основе int. Напишите функцию, которая принимает счет и возвращает сообщение о том, является ли он положительным, отрицательным или нулевым. */
type Score int
func sscore(scor1 Score) string {
	if scor1 == 0 {
		return "нулевым"
	} else if scor1 > 0 {
		return "положительным"
	} else  {
		return "отрицательным"
	}
}

/* 5.Классификация BMI
Определите тип BMI на основе float64. Напишите функцию, которая принимает значение BMI и возвращает сообщение о том, в какой категории оно находится: "Underweight", "Normal weight", "Overweight", или "Obesity". */
type BMI float64
func bm(bmi1 BMI) string {
	if bmi1 < 45  {
		return "Underweight"
	} else if bmi1 > 45 && bmi1 < 70 {
		return "Normal weight"
	} else if bmi1 > 70 && bmi1 > 90 {
		return "Overweight"
	} else {
		return "Obesity"
	}
}

/* 6.Возведение в квадрат
Определите тип функции UnaryOperation, которая принимает int и возвращает int. Напишите функцию для возведения числа в квадрат и используйте тип UnaryOperation для вызова этой функции. */
type UnaryOperation int 
func opeeration(unary UnaryOperation) UnaryOperation {
	return unary * unary 
}

/* 7.Удвоение числа
Определите тип функции UnaryOperation, которая принимает int и возвращает int. Напишите функцию для удвоения числа и используйте тип UnaryOperation для вызова этой функции.*/
func operationn(unar UnaryOperation) UnaryOperation {
	return unar * 2
}

/* 8.Проверка четности
Определите тип функции Check, которая принимает int и возвращает bool. Напишите функцию для проверки четности числа и используйте тип Check для вызова этой функции.*/
type Check int
func chekk(chek1 Check) bool {
	if chek1 % 2 == 0 {
		return true
	} else {
		return false 
	}
} 

/* 9.Проверка положительности
Определите тип функции Check, которая принимает int и возвращает bool. Напишите функцию для проверки, является ли число положительным, и используйте тип Check для вызова этой функции. */
func cheslo(chek2 Check) bool {
	if chek2 > 0 {
		return true
	} else {
		return false
	}
}

/* 10.Комбинирование функций
Определите тип функции Operation, которая принимает два int и возвращает int. Напишите функции для сложения, вычитания и умножения. Напишите функцию, которая принимает два int и массив функций Operation, и возвращает массив результатов применения этих функций. */
type Operationn int 
func plas (a, b Operationn) Operationn {
	return a + b 
}
func minas (a, b Operationn) Operationn {
	return a - b 
} 
func umnozh (a, b Operationn) Operationn {
	return a * b 
}

/* 11.Обратный отсчет
Создайте псевдоним для типа int под названием Count. Напишите функцию, которая принимает Count и выводит обратный отсчет до нуля.*/
type Count = int
func otshet (coun1 Count) {
	for i := coun1; i > 0; i --{
		fmt.Println(i)
	}
}

/* 12.Проверка уровня батареи
Создайте псевдоним для типа int под названием BatteryLevel. Напишите функцию, которая принимает BatteryLevel и возвращает сообщение о том, низкий, средний или высокий уровень заряда.*/
type BatteryLevel = int
func level(batter BatteryLevel) string {
	if batter == 1 && batter <= 20 {
		return "низкий уровень заряда"
	} else if batter > 20 && batter < 60 {
		return "средний уровень заряда"
	} else {
		return "высокий уровень заряда"
	}
}

/* 13.Определение категории веса
Создайте псевдоним для типа float64 под названием Weight. Напишите функцию, которая принимает Weight и возвращает сообщение о том, в какой категории веса находится объект: "Light", "Medium" или "Heavy".*/
type Weight = float64
func ves (a Weight) string {
	if a > 0 && a < 50 {
		return "Light"
	} else if a > 50 && a < 80 {
		return "Medium"
	} else {
		return "Heavy"
	}
}

/* 14.Оценка успеваемости
Создайте псевдоним для типа int под названием Grade. Напишите функцию, которая принимает Grade и возвращает сообщение о том, является ли оценка удовлетворительной (50 и выше) или нет.*/
type Grade = int
func osenka (b Grade) string {
	if b >= 50 {
		return "оценка удовлетворительной"
	} else {
		return "No"
	}
}

/* 15.Оценка состояния здоровья
Создайте псевдонимы для типов float64 под названиями BMI и BloodPressure. Напишите функцию, которая принимает BMI и BloodPressure, и возвращает сообщение о состоянии здоровья: "Healthy", "At risk" или "Unhealthy".*/
type BloodPressure = float64
func zdorovie (bmi BMI, bp BloodPressure )string {
	if bmi < 14 || bmi > 28 {
        if bp < 90 || bp > 140 {
            return "Unhealthy"
        }
        return "At risk"
    } else {
        if bp < 90 || bp > 140 {
            return "At risk"
        }
        return "Healthy"
    }
}

/* 16.Описание продукта
Создайте структуру Product с полями Name и Price. Напишите функцию, которая принимает Product и выводит его описание. */
type Product struct {
	Name string
	Price int
}
func produc(product1 Product) string {
	str := fmt.Sprintf("Имя: %s, цена: %d", product1.Name, product1.Price)
	return str
}

/* 17.Вывод информации о человеке
Создайте структуру Person с полями FirstName, LastName и Age. Напишите функцию, которая принимает Person и выводит его полное имя и возраст.*/
type Person struct {
	FirstName, LastName string
	Age int
}

func people(perssson Person) string {
	srt := fmt.Sprintf("Имя: %s, Фамиля: %s, Возраст: %d", perssson.FirstName, perssson.LastName, perssson.Age)
	return srt
}

/* 18.Сравнение продуктов
Создайте структуру Product с полями Name и Price. Напишите функцию, которая принимает два Product и возвращает более дорогой продукт.*/
func prodyct (apple, banana Product) string {
	if apple.Price < banana.Price {
		k := fmt.Sprintf("Имя: %s, цена: %d", banana.Name, banana.Price)	
		return k
	} else {
		k := fmt.Sprintf("Имя: %s, цена: %d", apple.Name, apple.Price)
		return k
	}
} 

/* 21.Обновление цены продукта
Создайте структуру Product с полями Name и Price. Напишите функцию, которая принимает указатель на Product и обновляет его цену.*/
func proddduct(b *Product, newPrice float64) {
	b.Price = int(newPrice)
} 

/* 22.Увеличение возраста
Создайте структуру Person с полями Name и Age. Напишите функцию, которая принимает указатель на Person и увеличивает его возраст на 1.*/
func personn (d *Person) {
	d.Age = d.Age + 1
}

/* 23.Обновление информации о продукте
Создайте структуру Product с полями Name и Price. Напишите функцию, которая принимает указатель на Product и обновляет его название и цену.*/
func ppproduct(f *Product, newname string, newprise int) {
	f.Name = newname
	f.Price = newprise
}

/* 24.Увеличение количества предметов
Создайте структуру Item с полями Name и Quantity. Напишите функцию, которая принимает указатель на Item и увеличивает количество на заданное значение.*/
type Item struct {
	Name string
	Quantity int
}
var h int 
func items(kolichestvo *Item, uvel int) {
	kolichestvo.Quantity += uvel
}

/* 26.Адрес человека
Создайте структуру Address с полями Street и City. Создайте структуру Person с полями Name и Address. Напишите функцию, которая принимает Person и выводит его имя и адрес.*/
type Address struct {
	Street, City string
}
type Persen struct {
Name string
Address Address
}
func adres(p Persen) {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Addres: %s, %s\n", p.Address.Street, p.Address.City)
}

/* 27.Описание компании
Создайте структуру Company с полями Name и Location (структура Address). Напишите функцию, которая принимает Company и выводит информацию о компании.*/
type Company struct {
	Name string
	Location Address
}
func company(c Company) {
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Addres: %s, %s\n", c.Location.Street, c.Location.City)
}

/* 28.Описание курса
Создайте структуру Course с полями Title и Instructor (структура Person). Напишите функцию, которая принимает Course и выводит информацию о курсе.*/
type Course struct {
	Title string
	Instructor Persen
}
func cours(s Course) {
	fmt.Printf("Title: %s\n", s.Title)
	fmt.Printf("Instructor: %s\n", s.Instructor.Name)
	fmt.Printf("Instructor Addres: %s, %s\n", s.Instructor.Address.Street, s.Instructor.Address.City)
}

/* 29.Описание события
Создайте структуру Event с полями Title и Location (структура Address). Напишите функцию, которая принимает Event и выводит информацию о событии.*/
type Event struct {
	Title string
	Location Address
}
func even (e Event) {
	fmt.Printf("Title: %s\n", e.Title)
	fmt.Printf("Location: %s, %s\n", e.Location.Street, e.Location.City)
}


/* 30.Управление проектом
Создайте структуру Project с полями Name и Manager (структура Person, где Person имеет поле Name и Address (структура Address)). Напишите функцию, которая принимает Project и выводит информацию о проекте и менеджере.*/
type Project struct{
	Name string
	Manager Persen
}
func project(pro Project){
	fmt.Printf("Name: %s\n", pro.Name)
	fmt.Printf("Manager: %s\n", pro.Manager.Name)
	fmt.Printf("Addres: %s, %s\n", pro.Manager.Address.Street, pro.Manager.Address.City)
}

