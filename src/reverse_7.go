package main

import (
	"fmt"
	"math"
	"strconv"
)

func main(){
	ret := reverse(123)
	fmt.Println(ret)
}

func reverse(x int) int {
	bs := []byte(strconv.FormatInt(int64(x), 10))
	newbs := make([]byte, len(bs))
	if x < 0 {
		for i := 1; i < len(bs); i ++ {
			newbs[len(newbs) - i] = bs[i]
		}
		newbs[0] = '-'
	} else {
		for i, b := range bs {
			newbs[len(newbs) - i - 1] = b
		}
	}
	ret, _ := strconv.ParseInt(string(newbs), 10, 64)
	if float64(ret) < -1 * math.Pow(2, 31) || float64(ret) > math.Pow(2, 31) - 1 {
		ret = 0
	}
	return int(ret)
}