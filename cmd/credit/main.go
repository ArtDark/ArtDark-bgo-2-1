package main

import (
	"fmt"
	"github.com/ArtDark/bgo_2_func/pkg/credit"
)

func main() {
	creditSum, period, percent := 1_000_000_00, 36, 20

	monthly, over, total := credit.Calculate(creditSum, period, percent)

	fmt.Println("Ежемесячный платеж: ", monthly)
	fmt.Println("Переплата: ", over)
	fmt.Println("Всего: ", total)

}
