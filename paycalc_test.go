package test

import (
	"fmt"
)

func main() {
	amt, err := paycalc.CalcPay(123.15)
	fmt.Println(amt)
	fmt.Println(err)
}
