package db

// The idea is that this database interface would serve all DB needs from
// the different service packages we implement. Then they could easily
// switch between different database implementations, without being
// specifically tied to one in particular, or even tied to the way other
// services are using their database implementations.

type Database interface {
	Get(table, id string) (interface{}, error)
	Set(table string, data interface{}) error
}
