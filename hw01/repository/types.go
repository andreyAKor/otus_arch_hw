package repository

// Интерфейс хранения данных
type Storager interface {
	Get(key string) interface{}
	Set(key string, value interface{})
}
