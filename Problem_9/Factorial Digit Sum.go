package Problem_9

import (
	"fmt"
	"math/big"
	"strconv"
)

func Factorial(n int) int {
	var bigInteger big.Int
	bigInteger.MulRange(1, int64(n))

	sum := 0                 //hold the sum of factorials
	x := bigInteger.String() //convert factorial result to string for sum each of it

	for _, item := range x {
		s, _ := strconv.Atoi(string(item)) //convert string to int
		sum += s                           //sum all number
	}

	fmt.Print(bigInteger.Mul(bigInteger, 1))
	return sum
}
