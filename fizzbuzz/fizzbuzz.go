package fizzbuzz

import (
	"strconv"
	"strings"
)

//Fizzbuzz must return a list of strings with numbers from 1 to limit, where:
// all multiples of int1 are replaced by str1, all multiples of int2 are
// replaced by str2, all multiples of int1 and int2 are replaced by str1str2
func Fizzbuzz(int1, int2, limit int, str1, str2 string) ([]string, error) {

	ret := make([]string, 0, limit-1)

	n := 1
	fizzBuzzWord := strings.Join([]string{str1, str2}, "")
	for n <= limit {
		if n%int1 == 0 && n%int2 == 0 {
			ret = append(ret, fizzBuzzWord)
		} else if n%int1 == 0 {
			ret = append(ret, str1)
		} else if n%int2 == 0 {
			ret = append(ret, str2)
		} else {
			ret = append(ret, strconv.Itoa(n))
		}
		n++
	}

	return ret, nil
}
