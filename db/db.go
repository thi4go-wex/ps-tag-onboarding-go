package db

type Database interface {
	Get(table, id string) (interface{}, error)
	Set(table string, data interface{}) error
}
