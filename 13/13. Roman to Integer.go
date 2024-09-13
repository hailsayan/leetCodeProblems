package main

import "fmt"

func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	length := len(s)

	for i := 0; i < length; i++ {
		if i < length-1 && romanMap[s[i]] < romanMap[s[i+1]] {
			total -= romanMap[s[i]]
		} else {
			total += romanMap[s[i]]
		}
	}

	return total
}
func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
