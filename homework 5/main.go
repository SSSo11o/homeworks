package main
import "fmt"
func main (){

	/*1.Учет накопительных счетов с ежемесячным пополнением
--Начальный баланс накопительного счета равен 100000 рублей.
--Реализуйте функции для пополнения счета каждый месяц на фиксированную сумму.
--Выводите баланс после каждого пополнения.
--Если баланс становится больше 500000 рублей, выведите сообщение "Достигнут лимит накоплений". */
balance := 100000
depozit := 4000
for {
	balance += depozit
	fmt.Println(balance)
	if balance >= 500000 {
	fmt.Println("Достигнут лимит накоплений")
	break
}
}

/*2.Рассчет выплат по кредиту с фиксированной ставкой
Начальная сумма кредита равна 3000000 рублей.
Реализуйте функции для ежемесячного расчета выплат по кредиту с фиксированной процентной ставкой 10%.
Выводите остаток по кредиту после каждой выплаты.
Если остаток по кредиту становится меньше 500000 рублей, выведите сообщение "Почти погашен кредит". */
kredit := 3000000
prosent := 300000
for {
	kredit -= prosent 
	fmt.Println(kredit) 
	if kredit < 500000 {
		fmt.Println("Почти погашен кредит")
		break
	}
}

/* Начальный баланс счета равен 500000 рублей.
Реализуйте функции для выполнения банковских переводов.
Если сумма перевода превышает 100000 рублей, выведите сообщение "Сумма перевода превышает лимит".
Выводите остаток на счете после каждого перевода. */
nachalbalance := 500000
viplota := 7000
for {
	nachalbalance -= viplota
	fmt.Println(nachalbalance)
	if nachalbalance < 100000 {
		fmt.Println("Сумма перевода превышает лимит")
		break
	}
}

/* 4.Учет процентов по вкладам с ежемесячной капитализацией
Начальный вклад равен 100000 рублей.
Реализуйте функции для ежемесячного начисления процентов по ставке 5% годовых.
Капитализируйте проценты ежемесячно и выводите сумму вклада после каждого начисления.
Если сумма вклада превышает 150000 рублей, выведите сообщение "Достигнут лимит вклада". */
vklad := 100000
prosentStavke := 5000
for {
	vklad += prosentStavke
	fmt.Println(vklad)
	if vklad >= 150000 {
		fmt.Println("Достигнут лимит вклада")
		break
	}
} 

/*6.Начисление сложных процентов на вклад
Начальная сумма вклада равна 200000 рублей.
Реализуйте функции для начисления сложных процентов каждый месяц по ставке 5%.
Выводите баланс после каждого начисления.
Если баланс становится больше 300000 рублей, выведите сообщение "Достигнут лимит начислений". */
nachalVklad := 200000
stavkaProsent := 10000
for {
	nachalVklad += stavkaProsent
	fmt.Println(nachalVklad)
	if nachalVklad > 300000 {
		fmt.Println("Достигнут лимит начислений")
		break
	}
}

addraskhod(1200)
addraskhod(45)

/* 8.Начисление процентов на депозит с ежегодной проверкой
Начальная сумма депозита равна 500000 рублей.
Реализуйте функции для начисления процентов каждый год по ставке 6%.
Выводите баланс после каждого начисления.
Если баланс становится больше 700000 рублей, выведите сообщение "Достигнут лимит начислений".*/
sumdepozit := 500000
prosentt := 30000
for {
	sumdepozit += prosentt 
	fmt.Println(sumdepozit)
	if sumdepozit > 700000 {
		fmt.Println("Достигнут лимит начислений")
		break
	}
}
}

/* 7.Учет ежедневных расходов с лимитом
Начальный лимит расходов в день равен 5000 рублей.
Реализуйте функции для добавления расходов в течение дня.
Выводите текущую сумму расходов после каждого добавления.
Если сумма расходов превышает лимит, выведите сообщение "Превышен дневной лимит". */
var limit float64 = 5000
var sumRaskhod float64 = 0
func addraskhod(a float64 ) {
	sumRaskhod += a
	fmt.Println(sumRaskhod)
	if sumRaskhod > limit {
		fmt.Println("Превышен дневной лимит")
	}
}


