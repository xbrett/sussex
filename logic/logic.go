package logic

import (
	"github.com/xbrett/sussex/database"
)

type Logic struct {
	mydb database.DatabaseAccess
}

//New creates a new logic pointer to the database layer
func New() Logic {
	//database.NewSQL() to use mySQL db
	//database.NewInMemoryDB() to use local db
	l := Logic{mydb: database.NewSQL()}
	return l
}
