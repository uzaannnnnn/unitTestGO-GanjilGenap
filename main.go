package main

import "fmt"

// IsEven mengembalikan true jika angka genap, false jika ganjil
func IsEven(num int) bool {
	return num%2 == 0
}

func main() {
	fmt.Println("5 adalah genap?", IsEven(5)) // Output: false
	fmt.Println("4 adalah genap?", IsEven(4)) // Output: true
}