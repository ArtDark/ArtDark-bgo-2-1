package credit


func pow(num, pow uint64) uint64 {
	var sum uint64 = num
	for i := 1; i < int(pow); i++ {
		sum = sum * num
	}
	return sum
}