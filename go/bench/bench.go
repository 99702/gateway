package bench

import (
	"fmt"
	"math/big"
	"time"
)

func Pprint[Out any](funcName string, fn func() Out) {
	start := time.Now()
	results := fn()
	elapsed := time.Since(start)

	switch v := any(results).(type) {
	case *big.Int:
		fmt.Printf("Function: %s |  Result: %s | Time: %s\n",
			funcName, bigIntToScientific(v), elapsed)
	case big.Int:
		fmt.Printf("Function: %s |  Result: %s | Time: %s\n",
			funcName, bigIntToScientific(&v), elapsed)
	default:
		fmt.Printf("Function: %s |  Result: %v | Time: %s\n",
			funcName, v, elapsed)
	}
}

func bigIntToScientific(n *big.Int) string {
	s := n.String()
	if len(s) <= 10 {
		return s
	}
	return fmt.Sprintf("%s.%se+%d", s[:1], s[1:6], len(s)-1)
}
