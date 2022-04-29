package main

import (
	"fmt"
	"math"
)

//剑指 Offer II 002. 二进制加法

func main() {
	//addBinary("10", "111")
	fmt.Println(addBinary("101111", "10"))
}

func addBinary(a string, b string) string {
	maxlength := int(math.Max(float64(len(a)), float64(len(b))))
	if len(a) < maxlength {
		x := make([]byte, maxlength)
		for i := 0;i < maxlength - len(a);i++ {
			x[i] = '0'
		}
		x = x[:maxlength - len(a)]
		for _, c := range a {
			x = append(x, byte(c))
		}
		a = string(x)
	}
	if len(b) < maxlength {
		x := make([]byte, maxlength)
		for i := 0;i < maxlength - len(b);i++ {
			x[i] = '0'
		}
		x = x[:maxlength - len(b)]
		for _, c := range b {
			x = append(x, byte(c))
		}
		b = string(x)
	}
	forward := false
	r := make([]byte, len(a) + 1)
	for i := maxlength - 1;i >= 0;i-- {
		if a[i] == '1' && b[i] == '1' {
			if forward {
				r[i+1] = '1'
			} else {
				r[i+1] = '0'
			}
			forward = true
		} else if a[i] == '0' && b[i] == '0' {
			if forward {
				r[i+1] = '1'
			} else {
				r[i+1] = '0'
			}
			forward = false
		} else {
			if forward {
				r[i+1] = '0'
				forward = true
			} else {
				r[i+1] = '1'
				forward = false
			}
		}
	}
	if forward {
		r[0] = '1'
		return string(r)
	} else {
		return string(r[1:])
	}
}