package algorithms

import (
	"fmt"
	"github.com/cornelk/hashmap"
	"testing"
)

func BenchmarkCanConstruct(b *testing.B) {
	result := new(WordResult)
	result.Memoizer = hashmap.New[string, bool]()
	for i := 0; i < 1; i++ {
		//result = CanConstruct("abcdef", &[]string{"ab", "abc", "cd", "def", "abcd"}, result)
		//fmt.Println(result.State)
		//result = CanConstruct("skateboard", &[]string{"bo", "rd", "ate", "t", "ska", "sk", "boat"}, result)
		//fmt.Println(result.State)
		//result = CanConstruct("enterapotentpot", &[]string{"a", "p", "ent", "enter", "ot", "o", "t"}, result)
		//fmt.Println(result.State)
		result = CanConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", &[]string{"e", "ee", "eee", "eeee", "eeeeee"}, result)
		fmt.Println(result.State)
	}
}
