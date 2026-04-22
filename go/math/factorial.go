package main

import (
	"gateway/bench"
	"math/big"
)

/*
factorial of n number: 1 * 2 * 3 * 4 ......................... +n

1: 1 = 1
2: 1*2 = 2
3: 1*2*3 = 6
4: 1*2*3*4 = 24
5: 1*2*3*4*5 = 120
6: 1*2*3*4*5*6 = 720
10: 1*2*3*4*5*6*7*8*9*10 = 3628800
11: 1*2*3*4*5*6*7*8*9*10*11 = 39916800
12: 1*2*3*4*5*6*7*8*9*10*11*12 = 479001600
13: 1*2*3*4*5*6*7*8*9*10*11*12*13 = 6227020800
500: 1*2*3*4*5*6*7*8*...... 500 = 6227020800

f(0) = 0 (n =1)
f(1) = 1 (n =1)
f(2) = f(1)*2 (n=2)
f(3) = f(2)*3 (n=2)
f(4) = f(3)*4 (n=4)
f(5) = f(4)*5 (n=5)
f(6) = f(5)*6 (n=6)
~~
f(k) = f(k-1)*k (n=k)
*/
func facto_of_n_number(n int) int {
	total := 1
	if n == 0 {
		return 1
	}
	for i := 1; i <= n; i++ {
		total *= i
	}
	return total
}

func big_int_facto_of_n_number(n int) *big.Int {
	total := big.NewInt(1)
	if n == 0 {
		return big.NewInt(1)
	}
	tempInt := new(big.Int)
	for i := 1; i <= n; i++ {
		tempInt.SetInt64(int64(i))
		total.Mul(total, tempInt)
	}
	return total
}

func recursion_facto_of_n_number(n int) int {
	if n > 0 {
		return recursion_facto_of_n_number(n-1) * n
	}
	return 1
}

func main() {
	const n = 12
	bench.Pprint(facto_of_n_number, 10)
	bench.Pprint(big_int_facto_of_n_number, 10)
	bench.Pprint(recursion_facto_of_n_number, 10)
}
