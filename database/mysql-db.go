// The mySQL database being used should have tables that match the following
//
//                  ----------GROCERIES----------
// +-------+-------------+------+-----+---------+----------------+
// | Field | Type        | Null | Key | Default | Extra          |
// +-------+-------------+------+-----+---------+----------------+
// | ID    | int(11)     | NO   | PRI | NULL    | auto_increment |
// | item  | varchar(50) | YES  |     | NULL    |                |
// +-------+-------------+------+-----+---------+----------------+

//                ----------Chores----------
// +-------+-------------+------+-----+---------+-------+
// | Field | Type        | Null | Key | Default | Extra |
// +-------+-------------+------+-----+---------+-------+
// | name  | varchar(20) | YES  |     | NULL    |       |
// | chore | varchar(20) | YES  |     | NULL    |       |
// +-------+-------------+------+-----+---------+-------+

package database

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	//drivers for mysql connection
	_ "github.com/go-sql-driver/mysql"
)

//SQLdb is a struct to restrict access to the db
type SQLdb struct {
	db *sql.DB
}

type ChorePair struct {
	Name  string
	Chore string
	ID    int
}

func openDatabaseConnection() *sql.DB {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/SussexSite")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func (s *SQLdb) insertIntoGroceries(item interface{}) error {
	return errors.New("stuff")
}

func (s *SQLdb) UpdateChores(item interface{}) error {
	var cp ChorePair
	cp = item.(ChorePair)
	_, err := s.db.Exec("UPDATE chores SET chore=? WHERE id=? VALUES(?, ?)", cp.Chore, cp.ID)
	return err
}

func (s *SQLdb) insertIntoMovies(item interface{}) error {
	return errors.New("movies error")
}

//Delete removes an entry based on id from the products table in productInfo db
func (s *SQLdb) Delete(id string, table string) error {
	_, err := s.db.Exec("DELETE FROM products WHERE id=?", id)
	return err
}

//Insert puts given product information into the products table in the db
// func (s *SQLdb) Insert(item interface{}, table string) error {
// 	var err error
// 	switch table {
// 	case "groceries":
// 		err = s.insertIntoGroceries(item)
// 	case "chores":
// 		err = s.insertIntoChores(item)
// 	case "movies":
// 		err = s.insertIntoMovies(item)
// 	default:
// 		err = errors.New("No table with name: " + table)
// 	}
// 	return err
// }

//Update changes the products quantity
func (s *SQLdb) Update(item interface{}, table string) error {
	_, err := s.db.Exec("UPDATE ? SET chore=? WHERE id=?", table)
	return err
}

//Get returns the product info for a given id
func (s *SQLdb) Get(id string, table string) string {
	queryStr := fmt.Sprintf("SELECT * FROM %s WHERE id=%s", table, id)
	res, err := s.buildJSON(queryStr)
	if res == "" || err != nil {
		res = "[]"
	}
	return res
}

//Print prints product information from database
func (s *SQLdb) Print(table string) string {
	res, _ := s.buildJSON("SELECT * FROM " + table)
	if res == "" {
		res = "[]"
	}
	return res
}

func (s *SQLdb) buildJSON(queryStr string) (string, error) {
	rows, err := s.db.Query(queryStr)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	//fmt.Println(string(jsonData))
	return string(jsonData), nil
}
