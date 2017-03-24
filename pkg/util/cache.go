package util

type cache struct {
	cache *map[string]cachedElement
}

type cachedElement struct {
	Words *[]string
	Count int
}

func (c *cache) Get(k string) *cachedElement {
	ca := *c.cache
	if val, ok := ca[k]; ok {
		return &val
	}
	return nil
}

func (c *cache) Put(k string, w *[]string, count int) {
	ce := new(cachedElement)
	ce.Words = w
	ce.Count = count
	ca := *c.cache
	ca[k] = *ce
}

var Cache *cache

func init() {
	c := make(map[string]cachedElement)
	Cache = new(cache)
	Cache.cache = &c
}
