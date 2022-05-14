package interfaces

import "time"

type CachedStorage interface {
	Get(interface{}) (string, error)
	Set(key, value string, expires time.Duration) error
	Del(key string) error
}
