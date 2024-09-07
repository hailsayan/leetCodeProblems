// package main

// import "fmt"

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversedHalf := 0

	for x > reversedHalf {
		reversedHalf = reversedHalf*10 + x%10
		x /= 10
	}

	return x == reversedHalf || x == reversedHalf/10
}

// func main() {
// 	fmt.Println(isPalindrome(121))
// 	fmt.Println(isPalindrome(-121))
// 	fmt.Println(isPalindrome(10))
// }