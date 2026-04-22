package main

import (
	"gateway/bench"
	"math/big"
)

/*
sum of n number: 1 + 2 + 3 + 4 ......................... +n

3: 1+2+3 = 6
4: 1+2+3+4 = 10
5: 1+2+3+4+5 = 15
10: 1+...+10 = 55
13: 1+...+13 = 91
500: 1+...+500 = 125,250

f(3) = f(2)+3 (n=2)
f(4) = f(3)+4 (n=4)
f(5) = f(4)+5 (n=5)
.....
f(10) = f(9)+6 (n=6)
~~
f(k) = f(k-1)+k (n=k)
*/

func sum_of_n_number1(n *big.Int) *big.Int {
	total := big.NewInt(0)
	tempInt := new(big.Int)
	one := big.NewInt(1)
	for i := big.NewInt(0); int(i.Int64()) <= int(n.Int64()); i.Add(i, one) {
		tempInt.Set(i)
		total.Add(total, tempInt)

	}
	return total
}
func recursive_sum_of_n_number(n *big.Int) *big.Int {
	if n.Sign() >= 0 {
		nMinusOne := new(big.Int).Sub(n, big.NewInt(1))
		return new(big.Int).Add(recursive_sum_of_n_number(nMinusOne), n)
	}
	return big.NewInt(1)
}
func main() {
	bench.Pprint(recursive_sum_of_n_number, big.NewInt(1000000))
	bench.Pprint(sum_of_n_number1, big.NewInt(1000000))
}
