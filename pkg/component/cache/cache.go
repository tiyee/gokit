package cache

type ICache interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
	Delete(key string) error
}

var _cache ICache

func InitCache(c ICache) {
	_cache = c
}
func Get(key string) ([]byte, error) {
	return _cache.Get(key)
}
func Set(key string, value []byte) error {
	return _cache.Set(key, value)
}

func Delete(key string) error {
	return _cache.Delete(key)
}
