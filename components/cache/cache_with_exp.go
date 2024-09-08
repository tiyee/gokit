package cache

import "time"

type ICacheE interface {
	ICache
	SetE(key string, value []byte, exp time.Duration) error
}
