package utils

import (
	"errors"
	"math"
)

// BeanToMoney 金豆转现金, 写死先除以10000，再乘以汇率！
func BeanToMoney(coin int) (m int, rate int) {
	r := 100
	return coin / 10000 * r, r
}

// MoneyToBean 现金转金豆
func MoneyToBean(money float64) (maxCoin int, ratio int, err error) {
	r := 100
	coin := money * 10000 / float64(r)
	c, b := math.Modf(coin)
	if b != 0 {
		return 0, 0, errors.New("you must enter an integer multiple of 100")
	}
	return int(c), r, nil
}

// float 保留n位小数
func Round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}
