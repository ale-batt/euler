package main

import "fmt"

func eratosthenes(limit int) []int {
	tab := make([]bool, limit+1)
	ret := make([]int, 0)

	for i := 2; i <= limit; i++ {
		if tab[i] == false {
			ret = append(ret, i)
			for j := i; j <= limit; j += i {
				tab[j] = true
			}
		}
	}
	return ret
}

func main() {
	var limit int = 2000000
	var primes []int
	var sum int

	primes = eratosthenes(limit)
	sum = 0
	for _, e := range primes {
		sum += e
	}
	fmt.Println("The sum of the", limit, "th primes is", sum)
}
