package main

import "fmt"

func fibonacci(n int) int {
	if n < 2 {
		return 1
	}
	res := fibonacci(n-1) + fibonacci(n-2)
	return res
}
func biggerNum(s ...int) int {
	var temp int
	for j := 0; j < len(s); j++ {
		for i := 0; i < len(s)-1; i++ {
			if i < len(s) && s[i] > s[i+1] {
				temp = s[i]
				s[i] = s[i+1]
				s[i+1] = temp
			}
		}
	}
	return s[len(s)-1]
}
func generadorImpares() func() uint {
	i := uint(1) // i permanecerá en el clousure de la función anónima a retornar
	return func() uint {
		var par = i
		i += 2
		return par
	}
}
func cambio(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	r := fibonacci(7)
	fmt.Println(r)

	s := []int{1, 5, 9, 15, 800, 25, 100, 200, 40, 10}
	bigger := biggerNum(s...)
	fmt.Println(bigger)

	nextPar := generadorImpares()
	fmt.Println(nextPar())
	fmt.Println(nextPar())
	fmt.Println(nextPar())

	a := 1
	b := 2
	fmt.Println(a, b)
	cambio(&a, &b)
	fmt.Println(a, b)

}
