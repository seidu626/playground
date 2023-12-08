package algorithms

import "fmt"

func StringReverse(str []byte, result []byte) []byte {
	if result == nil {
		result = make([]byte, 5)
	}

	if len(str) == 0 {
		return result
	}
	char := str[len(str) - 1]
	result = append(result, char)

	return StringReverse(str[:1], result)
}

func ReverseString(s []byte)  {
	if len(s) == 0 {
		return
	}
	ReverseString(s[1:])
	fmt.Printf("%s\n", string(s[0]))
	c := s[0]
	for i:=0; i < len(s) - 1; i++ {
		s[i] = s[i+1]
	}
	s[len(s) - 1] = c
}

func reverseString(input string) {
	if len(input) == 0 {
		return
	}
	reverseString(input[1:])
	fmt.Print(string(input[0]))
}


func SortedBinarySearch()  {

}