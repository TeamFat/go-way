package main

import (
	"fmt"
	"math"
	"strconv"
)

func StrToFloat64round(str string, prec int, round bool) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return FloatPrecision(f, prec, round)
}

// FloatPrecision float指定精度; round为true时, 表示支持四舍五入
func FloatPrecision(f float64, prec int, round bool) float64 {
	pow10N := math.Pow10(prec)
	if round {
		return math.Trunc((f+0.5/pow10N)*pow10N) / pow10N
	}
	return math.Trunc((f)*pow10N) / pow10N
}

func main() {
	f := StrToFloat64round("1.232325232", 5, true)
	fmt.Println(f)
}
