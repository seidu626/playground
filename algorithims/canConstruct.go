package algorithms

import (
	"github.com/cornelk/hashmap"
	"strings"
)

type WordResult struct {
	State    bool
	Words    *[]string
	Memoizer *hashmap.Map[string, bool]
}

func CanConstruct(target string, wordBank *[]string, result *WordResult) *WordResult {
	_, ok := result.Memoizer.Get(target)
	if ok {
		result.State = true
		return result
	}
	if target == "" {
		result.State = true
		return result
	}
	result.State = false

	for _, word := range *wordBank {
		if strings.HasPrefix(target, word) {
			suffix := strings.TrimPrefix(target, word)
			result = CanConstruct(suffix, wordBank, result)
			if result.State {
				result.Memoizer.Set(target, true)
				return result
			}
		}
	}

	return result

}
