package cache

type cacheController interface {
	GetData() string
	SendData(data string)
}

type dbCache struct {
	data string
}

// GetData returns data from cache
func (c *dbCache) GetData() string {
	return c.data
}

// SendData updates data in cache
func (c *dbCache) SendData(data string) {
	c.data = data
}

// NewCacheController returns new instance of cache
func NewCacheController() cacheController {
	return &dbCache{}
}
