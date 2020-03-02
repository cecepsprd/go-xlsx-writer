package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/subosito/gotenv"
	"github.com/tealeg/xlsx"
	"gitlab.twprisma.com/helper_db_lmd/helper"
)

func main() {
	err := gotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}

	db_user := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")

	fmt.Println(db_password)
	db, err := sqlx.Connect("mysql", db_user+":"+db_password+"@("+db_host+":3306)/LMD")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// helper.SelectAll()

	xlsFile, _ := xlsx.OpenFile("./lmd.xlsx")

	var wg sync.WaitGroup

	for _, sheet := range xlsFile.Sheets {
		switch sheet.Name {
		case "usergroup":
			wg.Add(1)
			go helper.InsertIntoUsergroup(sheet.Rows, &wg, db)
		case "user":
			wg.Add(1)
			go helper.InsertIntoUser(sheet.Rows, &wg, db)
		case "kk":
			wg.Add(1)
			go helper.InsertIntoKK(sheet.Rows, &wg, db)
		case "poskora":
			wg.Add(1)
			go helper.InsertIntoPoskora(sheet.Rows, &wg, db)
		}
	}
	wg.Wait()
}
