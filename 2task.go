package main

import "fmt"

func main() {
	var dep, per float64

	dep = inputDeposit()
	per = inputPercent()
	fmt.Println("Годовой процент: ",per, "\nСумма вклада: ", dep)
	period := 5

	for period>0{
		dep = riseDepositeInYear(dep, per)
		period--
	}

	fmt.Printf("За 5 лет депозит выростет до %.2f", dep)
}

func inputDeposit() float64{
	var a float64

	fmt.Println("Введите сумму вклада: ")
	fmt.Scan(&a)

	return a
}

func inputPercent() float64{
	var a float64

	fmt.Println("Введите процент: ")
	fmt.Scan(&a)

	return a
}

func riseDepositeInYear(dep, per float64) (newDep float64){
	newDep = dep*(1+per/100)
	return newDep
}