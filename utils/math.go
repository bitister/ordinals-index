package utils

import (
	"strings"

	decimal "github.com/shopspring/decimal"
)

func Sub(a float64, bs ...float64) float64 {
	decimala := decimal.NewFromFloat(a)
	for _, b := range bs {
		decimalb := decimal.NewFromFloat(b)
		decimala = decimala.Sub(decimalb)
	}
	f, _ := decimala.Float64()
	return f
}

func Add(a float64, bs ...float64) float64 {
	decimala := decimal.NewFromFloat(a)
	for _, b := range bs {
		decimalb := decimal.NewFromFloat(b)
		decimala = decimala.Add(decimalb)
	}
	f, _ := decimala.Float64()
	return f
}

func Mul(a float64, bs ...float64) float64 {
	decimala := decimal.NewFromFloat(a)
	for _, b := range bs {
		decimalb := decimal.NewFromFloat(b)
		decimala = decimala.Mul(decimalb)
	}
	f, _ := decimala.Float64()
	return f
}

func Div(a float64, bs ...float64) float64 {
	decimala := decimal.NewFromFloat(a)
	for _, b := range bs {
		decimalb := decimal.NewFromFloat(b)
		decimala = decimala.Div(decimalb)
	}
	f, _ := decimala.Float64()
	return f
}

//string转float保留指定精度
func String2Float(a string, decm int32) float64 {
	d, _ := decimal.NewFromString(a)
	f, _ := d.Truncate(decm).Float64()
	return f
}

//float保留指定精度
func Float2Float(a float64, decm int32) float64 {
	d := decimal.NewFromFloat(a)
	f, _ := d.Round(decm).Float64()
	return f
}

//精度处理
// ("100000000",8) = "10.00000000"
func Str2FloatStr(f string, m int) string {
	i := len(f)
	if i <= m {
		for i < m {
			f = "0" + f
			i++
		}
		return "0." + f
	} else {
		sliceTep := make([]string, i)
		for index, value := range f {
			if index == i-m {
				sliceTep = append(sliceTep, ".")
			}
			sliceTep = append(sliceTep, string(value))
		}
		return strings.Join(sliceTep, "")
	}
}

//精度处理
// ("1000000000",6,8) = 1000.00000000
func Str2Float(f string, m int, t int32) float64 {
	i := len(f)
	if i <= m {
		for i < m {
			f = "0" + f
			i++
		}
		d, _ := decimal.NewFromString("0." + f)
		f, _ := d.Float64()
		return f
	} else {
		sliceTep := make([]string, i)
		for index, value := range f {
			if index == i-m {
				sliceTep = append(sliceTep, ".")
			}
			sliceTep = append(sliceTep, string(value))
		}
		d, _ := decimal.NewFromString(strings.Join(sliceTep, ""))
		f, _ := d.RoundFloor(t).Float64()
		return f
	}

}

//精度处理
// ("100.000000",8) = 10000000000
func Float2String(value float64, d int) string {
	f := decimal.NewFromFloat(value)
	d1 := decimal.NewFromInt(10).Pow(decimal.NewFromInt(int64(d)))
	f2 := f.Mul(d1).String()
	if strings.Contains(f2, ".") {
		f2 = strings.Split(f2, ".")[0]
	}
	return f2
}
