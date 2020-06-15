package credit

import "math"

func Calculate(credit, period, percent int) (monthly, over, total int) {
	monthlyPercent := float64(percent) / 12 / 100

	exp := math.Pow(1+monthlyPercent, float64(period))
	coefficient := (monthlyPercent * exp) / (exp - 1)

	monthly = int(coefficient * float64(credit))
	over = period*monthly - credit
	total = monthly * period

	return

}