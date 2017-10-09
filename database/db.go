package database

type DatabaseAccess interface {
	//Insert(interface{}, string) error
	Get(string, string) string
	Print(string) string
	Update(interface{}, string) error
	Delete(string, string) error
	UpdateChores(interface{}) error
}

//NewSQL initializes a pointer to a sql database
func NewSQL() DatabaseAccess {
	sql := SQLdb{db: openDatabaseConnection()}
	return &sql
}
