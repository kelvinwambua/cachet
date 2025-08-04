package store

type Store interface {
	Get(key string) (string, bool)
	Set(key string, value string)
	Delete(key string) bool
	Exists(key string) bool
	Keys() []string
	Size() int
	Clear()
}
