package credit

import "math"

func Calculate(creditSum, period, percent uint64) (monthlyPay, overpayment, totalPay uint64) {
	monthlyPercent := float64(percent) / 12 / 100
	exponent := math.Pow((1 + monthlyPercent), float64(period))
	coefficient := monthlyPercent * exponent / (exponent - 1)

	monthlyPay = uint64(coefficient * float64(creditSum))
	overpayment = period*monthlyPay - creditSum
	totalPay = monthlyPay * period

	return

}
