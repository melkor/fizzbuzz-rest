package hit

import (
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	c := &Cache{
		client: redis.NewClient(&redis.Options{
			Addr: mr.Addr(),
		}),
	}

	int1 := 3
	int2 := 5
	limit := 20
	str1 := "fizz"
	str2 := "fuzz"

	expectedMostFrequentedRequest := makeKey(int1, int2, limit, str1, str2)

	for i := 0; i < 2; i++ {
		_, err := c.Add(int1, int2, limit, str1, str2)
		assert.NoError(t, err)
	}
	_, err = c.Add(1, 2, 3, "puzzle", "buble")
	assert.NoError(t, err)

	val, err := c.GetMostFrequentRequest()
	if err != nil {
		panic(err)
	}
	assert.Equal(t, expectedMostFrequentedRequest, val, "")

}
