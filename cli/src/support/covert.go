package support

import (
	"fmt"
	"math"
	"strings"
)

func ConvertStringToFloat(s string) float64 {
	var f float64

	fmt.Sscanf(strings.TrimSpace(s), "%f", &f)

	return Round(f)
}

func Round(amount float64) float64 {
	return math.Round(amount*100) / 100
}

func ConvertAmountToFloat(amount string) float64 {
	amount = strings.Replace(amount, "R$", "", -1)
	amount = strings.Replace(amount, ",", ".", -1)

	return ConvertStringToFloat(amount)
}
