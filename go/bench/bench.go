package bench

import (
	"fmt"
	"math/big"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func Pprint[In any, Out any](fn func(In) Out, arg1 In) {
	ptr := reflect.ValueOf(fn).Pointer()
	funcName := runtime.FuncForPC(ptr).Name()
	funcName = strings.ReplaceAll(funcName, "main.", "")

	start := time.Now()
	results := fn(arg1)
	elapsed := time.Since(start)

	switch v := any(results).(type) {
	case *big.Int:
		fmt.Printf("Function: %s | Input: [%v] | Result: %s | Time: %s\n",
			funcName, arg1, bigIntToScientific(v), elapsed)
	case big.Int:
		fmt.Printf("Function: %s | Input: [%v] | Result: %s | Time: %s\n",
			funcName, arg1, bigIntToScientific(&v), elapsed)
	default:
		fmt.Printf("Function: %s | Input: [%v] | Result: %v | Time: %s\n",
			funcName, arg1, v, elapsed)
	}
}

func bigIntToScientific(n *big.Int) string {
	s := n.String()
	if len(s) <= 10 {
		return s
	}
	return fmt.Sprintf("%s.%se+%d", s[:1], s[1:6], len(s)-1)
}
