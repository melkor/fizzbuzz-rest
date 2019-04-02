package hit

import (
	"net/url"
	"strconv"
)

//Handler is an interface that provides 2 functions
// - Add to add a hit into historic
// - GetMostFrequentRequest to get the most frequent request
type Handler interface {
	Add(int1, int2, limit int, str1, str2 string) (float64, error)
	GetMostFrequentRequest() (string, error)
}

//make a unique key with request paramters
func makeKey(int1, int2, limit int, str1, str2 string) string {
	q := url.Values{}
	q.Add("int1", strconv.Itoa(int1))
	q.Add("int2", strconv.Itoa(int2))
	q.Add("limit", strconv.Itoa(limit))
	q.Add("str1", str1)
	q.Add("str2", str2)
	return q.Encode()
}
