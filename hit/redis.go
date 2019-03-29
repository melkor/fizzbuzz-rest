package hit

import "github.com/go-redis/redis"

//ZSetKey is the key name of ordered set
const ZSetKey = "fizzbuzz-hit"

//Cache is used to store hit into a Redis
type Cache struct {
	client *redis.Client
}

//NewCache initialize a Cache hti object
func NewCache(addr, password string, db int) *Cache {

	c := &Cache{
		client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		}),
	}
	return c
}

func (c *Cache) Add(int1, int2, limit int, str1, str2 string) (float64, error) {
	member := makeKey(int1, int2, limit, str1, str2)
	return c.client.ZIncr(
		ZSetKey,
		redis.Z{
			Score:  1,
			Member: member,
		},
	).Result()
}

//GetMostFrequentRequest return the most frequebt request
func (c *Cache) GetMostFrequentRequest() (string, error) {
	vals, err := c.client.ZRevRangeByScore(
		ZSetKey,
		redis.ZRangeBy{
			Max:    "+inf",
			Min:    "-inf",
			Offset: 0,
			Count:  1,
		},
	).Result()

	if err != nil {
		return "", err
	}

	if len(vals) > 0 {
		return vals[0], nil
	}

	return "", nil
}
