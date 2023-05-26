package cache

import "time"

type ICacheE interface {
	ICache
	SetE(key string, value []byte, exp time.Duration) error
}

var _cacheE ICacheE

func InitCacheE(c ICacheE) {
	_cacheE = c
}
func SetE(key string, value []byte, exp time.Duration) error {
	return _cacheE.SetE(key, value, exp)
}
