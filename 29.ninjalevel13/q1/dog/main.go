package dog

// Years ...converts human to dog years
func Years(y int) int {
	return y * 7
}

// YearsTwo ...human years to dog years by addition; slow
func YearsTwo(y int) int {
	count := 0
	for i:=0; i<y; i++ {
		count += 7
	}
	return count
}