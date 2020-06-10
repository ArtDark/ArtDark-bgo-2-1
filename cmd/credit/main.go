package main

import (
	"fmt"
	"netology/bgo_func/pkg/credit"
)

func main() {
	var creditSum uint64 = 100_000_00
	var period uint64 = 12
	var percent uint64 = 10

	monthlyPay, overpaymen, total := credit.Calculate(creditSum, period, percent)

	fmt.Println("Ежемесячный платеж: ", monthlyPay)
	fmt.Println("Переплата: ", overpaymen)
	fmt.Println("Всего: ", total)

}
