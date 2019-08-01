package main

import "fmt"

func main(){
	fmt.Println(romanToInt("XIV"))
}

func romanToInt(s string) int {
	m := map[string]int{
		"M": 1000, "CM": 900, "D": 500, "CD": 400, "C": 100, "XC": 90, "L": 50, "XL": 40, "X": 10, "IX": 9, "V": 5, "IV": 4, "I": 1,
	}
	var ret int
	if len(s) == 1 {
		return m[s]
	}
	for i :=0; i < len(s) -1; i++ {
		v := m[string(s[i])]
		if v < m[string(s[i + 1])] {
			v = -v
		}
		ret += v
	}
	ret += m[string(s[len(s) - 1])]
	return ret
}
