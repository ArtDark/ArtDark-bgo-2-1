package credit

import "math"


func Calculate(creditSum, period, percent int) (monthlyPay, overpayment, totalPay int) {
	monthlyPercent := math.Round((float64(percent)/12/100)*1000) / 1000 // V 166

	exp := math.Pow(1+monthlyPercent, float64(period)) // 1,834654562
	coefficient := (monthlyPercent * exp) / (exp - 1)

	monthlyPay = int(math.Round(coefficient*1_000_000)) * creditSum / 1_000_000
	overpayment = period*monthlyPay - creditSum
	totalPay = monthlyPay * period

	return

}
