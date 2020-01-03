package paycalc

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// CalcPay returns net pay based on grossPay
func CalcPay(grossPay float64) (float64, error) {
	err := fmt.Errorf("Gross Pay is %v", grossPay)
	logrus.Fatal(err)
	netPay := grossPay
	return netPay, err

}
