

//IMPROVEMENTS
// Let the user pick the PostGres schema
// Let the user specify generated location
// Let the user pick the package name


package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

type Datarows struct {
	tablename   string
	columnnames []string
	datatypes   []string
}

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "qwert0612"
	DB_NAME     = "postgres"
)

var TypeMap = map[string]string{
	"character varying" : "string",
	"integer" : "int",
}

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	check(err)
	
	rows, err := db.Query("select * from pg_catalog.pg_tables where schemaname = 'public'")
	check(err)
	defer rows.Close()
	
	columns, _ := rows.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	
	for rows.Next(){
		for i, _ := range columns{
			valuePtrs[i] = &values[i]
		}
		
		rows.Scan(valuePtrs...)
		tableName := string(values[1].([]byte))
		GenerateTableStruct(db, tableName)
		
	}
	
}

func GenerateTableStruct(db *sql.DB, tableName string) {
	f, err := os.Create(tableName + ".go")
	check(err)
	defer f.Close()

	rows, err := db.Query("SELECT column_name, data_type FROM information_schema.COLUMNS where TABLE_NAME = '" + tableName + "'")
	check(err)
	defer rows.Close()

	r := Datarows{}
	r.tablename = tableName
	for rows.Next() {
		var columnname string
		var datatype string
		rows.Scan(&columnname, &datatype)
		r.columnnames = append(r.columnnames, columnname)
		r.datatypes = append(r.datatypes, datatype)
	}

	GenerateHeader(f)
	GenerateStruct(f, r)
	GenerateInsert(f, r)
	GenerateGetById(f, r)
	GenerateUpdate(f, r)
	GenerateDelete(f, r)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}

}
