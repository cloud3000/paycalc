package paycalc

import "fmt"

func CalcPay(amt float64) (float64, Error) {
	err2 := fmt.Errorf("math: square root of negative number %v", amt)
	return amt, err2

}
