package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/wendal/go-oci8"
)

func query() {
	os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8")
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	db, err := sql.Open("oci8", "user/pwd@server")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("select 1 from dual")
	if err != nil {
		log.Fatal(err)
	}
	cols, _ := rows.Columns()
	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))
	dest := make([]interface{}, len(cols))
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}
	for rows.Next() {
		err = rows.Scan(dest...)
		for i, raw := range rawResult {
			if raw == nil {
				result[i] = ""
			} else {
				result[i] = string(raw)
			}
		}
		fmt.Printf("%s\n", result[0])
	}
	rows.Close()
}

func main() {
	query()
	time.Sleep(5 * time.Second)
}
