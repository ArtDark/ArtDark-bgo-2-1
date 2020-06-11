package credit

func Calculate(creditSum, period, percent uint64) (monthlyPay, overpayment, totalPay uint64) {
	monthlyPercent := percent * 100 / 12 // *divine 10_000

	coefficient := (monthlyPercent * pow(10_000 + monthlyPercent, period)) / (pow(10_000 + monthlyPercent, period) - 10_000)

	monthlyPay = coefficient * creditSum
	overpayment = period*monthlyPay - creditSum
	totalPay = monthlyPay * period

	return

}


func pow(num, pow uint64) uint64 {
	var sum uint64 = num
	for i := 1; i < int(pow); i++ {
		sum = sum * num
	}
	return sum
}