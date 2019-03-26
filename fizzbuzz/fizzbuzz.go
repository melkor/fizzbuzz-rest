package fizzbuzz

import (
	"errors"
	"strconv"
)

//Fizzbuzz must return a list of strings with numbers from 1 to limit, where:
// all multiples of int1 are replaced by str1, all multiples of int2 are
// replaced by str2, all multiples of int1 and int2 are replaced by str1str2
func Fizzbuzz(int1, int2, limit int, str1, str2 string) ([]string, error) {

	if int1 < 1 {
		return []string{}, errors.New("int1 can't be minus than 1")
	}

	if int2 < 1 {
		return []string{}, errors.New("int2 can't be minus than 1")
	}

	if limit < 1 {
		return []string{}, errors.New("limit can't be minus than 1")
	}

	ret := make([]string, 0, limit-1)

	n := 1
	for n <= limit {
		if n%int1 == 0 && n%int2 == 0 {
			ret = append(ret, str1+str2)
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
