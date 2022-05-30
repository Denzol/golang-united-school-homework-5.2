package cache

import "time"

type Cache struct {
	myMap   map[string]string
	flag    bool
	timeMap map[string]time.Time
}

func NewCache() Cache {
	var new Cache
	new.myMap = make(map[string]string)
	new.timeMap = make(map[string]time.Time)
	return new
}

func (c *Cache) Get(key string) (string, bool) {
	value, ok := c.myMap[key]
	if ok == true {
		c.flag = ok
		return value, c.flag
	} else {
		c.flag = ok
		return "", c.flag
	}
}

func (c *Cache) Put(key, value string) {
	c.myMap[key] = value
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.myMap[key] = value
	c.timeMap[key] = deadline
}

func (c *Cache) Keys() []string {
	u := time.Now()
	keys := []string{}
	for i, t := range c.timeMap {
		if u.After(t) {
			delete(c.myMap, i)
		}
	}
	for i, _ := range c.myMap {
		keys = append(keys, i)
	}
	return keys
}
