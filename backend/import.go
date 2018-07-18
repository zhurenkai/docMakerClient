package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)

func main()  {
	db,err := sql.Open(`mysql`,`root:123@tcp(127.0.0.1:3306)/information_schema?charset=utf8`)
	if err != nil {
		log.Print(err,`connect db faild`)
		return
	}
	rows,err := db.Query(`SELECT COLUMN_NAME,DATA_TYPE,COLUMN_COMMENT FROM COLUMNS WHERE COLUMN_COMMENT !="" ORDER BY COLUMN_COMMENT DESC`)
	//rows,err := db.Query(`SELECT * FROM Account limit 3`)
	if err != nil {
		log.Print(err,`exec failed`)
		return
	}
	for rows.Next() {
		var COLUMN_NAME,DATA_TYPE,COLUMN_COMMENT string
		rows.Scan(&COLUMN_NAME,&DATA_TYPE,&COLUMN_COMMENT)
		fmt.Println(COLUMN_NAME,DATA_TYPE,COLUMN_COMMENT)
	}
}