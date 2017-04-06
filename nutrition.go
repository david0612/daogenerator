package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"bytes"
	"strconv"
)

type nutrition struct {
	Id int
	Name string
	ServingSize string
	Calories int
	GramsFat int
	GramsCarbs int
	GramsProtein int
}

func (t *nutrition) Insert(db *sql.DB) {
	
	var id int
	
	buf := new(bytes.Buffer)
	buf.WriteString("\r\n")
	buf.WriteString("INSERT INTO public.\"nutrition\" (")
	buf.WriteString("\"Name\",")
	buf.WriteString("\"ServingSize\",")
	buf.WriteString("\"Calories\",")
	buf.WriteString("\"GramsFat\",")
	buf.WriteString("\"GramsCarbs\",")
	buf.WriteString("\"GramsProtein\"")
	buf.WriteString(") values (")
	buf.WriteString("$1,")
	buf.WriteString("$2,")
	buf.WriteString("$3,")
	buf.WriteString("$4,")
	buf.WriteString("$5,")
	buf.WriteString("$6)")
	buf.WriteString(" returning \"Id\";")
	
	err := db.QueryRow(buf.String(), t.Name, t.ServingSize, t.Calories, t.GramsFat, t.GramsCarbs, t.GramsProtein).Scan(&id)
	check(err)
	t.Id = id
}

func (t *nutrition) GetById(id int, db *sql.DB) {
	
	rows, err := db.Query("SELECT * FROM public.\"nutrition\" WHERE \"Id\" = " + strconv.Itoa(id))
	defer rows.Close()
	check(err)
	
	for rows.Next() {
		rows.Scan(&t.Id, &t.Name, &t.ServingSize, &t.Calories, &t.GramsFat, &t.GramsCarbs, &t.GramsProtein)
	}
}

func (t *nutrition) Update(db *sql.DB) (int64, error) {
	buf := new(bytes.Buffer)
	buf.WriteString("UPDATE public.\"nutrition\" SET ")
	buf.WriteString(" \"Name\"=$1,")
	buf.WriteString(" \"ServingSize\"=$2,")
	buf.WriteString(" \"Calories\"=$3,")
	buf.WriteString(" \"GramsFat\"=$4,")
	buf.WriteString(" \"GramsCarbs\"=$5,")
	buf.WriteString(" \"GramsProtein\"=$6")
	buf.WriteString(" WHERE \"Id\"=$7")
	
	stmt, err := db.Prepare(buf.String())
	check(err)
	
	res, err := stmt.Exec(t.Name, t.ServingSize, t.Calories, t.GramsFat, t.GramsCarbs, t.GramsProtein, t.Id)
	check(err)
	
	return res.RowsAffected()
}

func (t *nutrition) Delete(db *sql.DB) (int64, error) {
	
	stmt, err := db.Prepare("DELETE FROM public.\"nutrition\" WHERE \"Id\" = $1")
	check(err)
	
	res, err := stmt.Exec(t.Id)
	check(err)
	
	return res.RowsAffected()
}