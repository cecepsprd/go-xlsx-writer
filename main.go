package main

import (
	"sync"

	"github.com/cecepsprd/go-xlsx-writer/config"
	"github.com/cecepsprd/go-xlsx-writer/helper"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tealeg/xlsx"
)

func main() {

	db := config.DBConnect()
	defer db.Close()

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
