package hit

import (
	"net/url"
	"strconv"
)

//Hit represents an object store hit in memory
type Hit struct {
	higherHit           float64
	mostFrequentRequest string
	requests            map[string]float64
}

type HiteHandler interface {
	Add(int1, int2, limit int, str1, str2 string) (float64, error)
	GetMostFrequentRequest() (string, error)
}

func makeKey(int1, int2, limit int, str1, str2 string) string {
	q := url.Values{}
	q.Add("int1", strconv.Itoa(int1))
	q.Add("int2", strconv.Itoa(int2))
	q.Add("limit", strconv.Itoa(limit))
	q.Add("str1", str1)
	q.Add("str2", str2)
	return q.Encode()
}
