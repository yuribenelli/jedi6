package repo

func Foo() int {
	return 12
}
func Bar() (int, string) {
	return 37, "ciao"
}

func Somma(ii ...int) int {
	sum := 0
	for _, v := range ii {
		sum += v
	}
	return sum
}

func Somma2(ii []int) int {
	totale := 0
	for _, v := range ii {
		totale += v
	}
	return totale
}

func DefSum(a int, b int) int {
	return a + b
}
