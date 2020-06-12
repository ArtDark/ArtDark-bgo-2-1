package main

import (
	"fmt"
	"netology/bgo_func/pkg/credit"
)

func main() {
	var creditSum, period, percent int = 1_000_000_00, 36, 20

	var monthly, over, total int = credit.Calculate(creditSum, period, percent)

	fmt.Println("Ежемесячный платеж: ", monthly)
	fmt.Println("Переплата: ", over)
	fmt.Println("Всего: ", total)

}
