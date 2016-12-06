package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

func main(){
	db, err := sql.Open("mysql", "root:123456@/test?charset=utf8")
	checkErr(err)
	defer db.Close()

	//insert
	stmt, err := db.Prepare("Insert person SET name=?,age=?,sex=?")
	checkErr(err)

	res, err := stmt.Exec("doge", 18, "m")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	//update
	stmt, err = db.Prepare("update person set name=? where id=?")
	checkErr(err)

	res, err = stmt.Exec("dog", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//select
	rows, err := db.Query("SELECT * FROM person")
	checkErr(err)

	for rows.Next() {
		id, name, sex, age := 0, "", "", 0
		err = rows.Scan(&id, &name, &sex, &age)
		checkErr(err)
		fmt.Printf("id: %d, name: %s, sex: %s, age:%d\n", id, name, sex, age)
	}

	//delete
	stmt, err = db.Prepare("delete from person where id=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func checkErr(err error){
	if err != nil {
		panic(err)
	}
}
