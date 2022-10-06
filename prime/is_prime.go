package prime

import (
	"fmt"
	"math"
)

func IsPrimeExample() {
	ret := IsPrime(24)
	fmt.Println(ret)
}

func IsPrime(n int) bool {
	if n <= 2 {
		return true
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
