package algorithms

import "fmt"

//Given two non-negative integers, num1 and num2 represented as string, return the sum of num1 and num2 as a string.
//You must solve the problem without using any built-in library for handling large integers (such as BigInteger).
//You must also not convert the inputs to integers directly.

func AddStrings(num1 string, num2 string) string {
	num1Byte := []byte(num1)
	//num2Byte := []byte(num2)
	for i := 0; i <= len(num1Byte) - 1; i++ {
		fmt.Printf("Number: %d\n", num1Byte[i])
	}
	return ""
}