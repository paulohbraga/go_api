package lib

func Fact(numValues int) int {
	fat := 1
	for i := 1; i <= numValues; i++ {
		fat *= i
	}
	return fat
}

func FactRecursive(numvalues int) int {
	// slower than the iterative version

	if numvalues == 0 {
		return 1
	}
	if numvalues < 0 {
		return -1
	}
	return (numvalues * FactRecursive(numvalues-1))
}

func Permute(numvalues int, r int) int {
	return Fact(numvalues) / (numvalues - r)
}
