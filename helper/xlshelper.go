package helper

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/tealeg/xlsx"
	"gitlab.twprisma.com/helper_db_lmd/model"
)

func InsertIntoUsergroup(rows []*xlsx.Row, wg *sync.WaitGroup, db *sqlx.DB) {
	defer wg.Done()
	for i, row := range rows {
		if i > 0 {
			id, _ := strconv.Atoi(row.Cells[0].String())
			tx := db.MustBegin()
			_, err := tx.NamedExec("INSERT INTO usergroup (id,usergroup_name,note) VALUES (:id,:usergroup_name,:note)",
				&model.Usergroup{
					ID:            id,
					UsergroupName: row.Cells[1].String(),
					Note:          row.Cells[2].String(),
				})
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println("[usergroup]success")

			errorCommit := tx.Commit()
			if errorCommit != nil {
				log.Fatalln(errorCommit)
			}
		}
		continue
	}
}

func InsertIntoPoskora(rows []*xlsx.Row, wg *sync.WaitGroup, db *sqlx.DB) {
	defer wg.Done()

	for i, row := range rows {
		if i > 0 {
			tx := db.MustBegin()
			_, err := tx.NamedExec("INSERT INTO poskora(id,user_id,warehouse_id,code,is_active,max_member,type,shipping_address,village_id,latitude,longitude,is_wow,created_by) VALUES(:id,:user_id,:warehouse_id,:code,:is_active,:max_member,:type,:shipping_address,:village_id,:latitude,:longitude,:is_wow,:created_by)",
				&model.Poskora{
					ID:              toInt(row.Cells[0]),
					UserID:          toInt(row.Cells[1]),
					WarehouseID:     toInt(row.Cells[2]),
					Code:            row.Cells[3].String(),
					IsActive:        toInt(row.Cells[4]),
					MaxMember:       toInt(row.Cells[5]),
					Type:            row.Cells[6].String(),
					ShippingAddress: row.Cells[7].String(),
					VillageID:       toInt(row.Cells[8]),
					Latitude:        toFloat(row.Cells[9]),
					Longitude:       toFloat(row.Cells[10]),
					IsWow:           toInt(row.Cells[11]),
					CreatedBy:       toInt(row.Cells[12]),
				})
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println("[poskora]success")

			errorCommit := tx.Commit()
			if errorCommit != nil {
				log.Fatalln(errorCommit)
			}
		}
		continue
	}
}

func InsertIntoKK(rows []*xlsx.Row, wg *sync.WaitGroup, db *sqlx.DB) {
	defer wg.Done()
	for i, row := range rows {
		if i > 0 {
			id, _ := strconv.Atoi(row.Cells[0].String())

			tx := db.MustBegin()
			_, err := tx.NamedExec("INSERT INTO kk(id,poskora_id,name,address,village_id,phone,is_deleted,is_active,is_verified,is_poskora,is_default,saldo_balance,total_order,created_by,updated_by,survey_answered) VALUES (:id,:poskora_id,:name,:address,:village_id,:phone,:is_deleted,:is_active,:is_verified,:is_poskora,:is_default,:saldo_balance,:total_order,:created_by,:updated_by,:survey_answered)",
				&model.KK{
					ID:             id,
					PoskoraID:      toInt(row.Cells[1]),
					Name:           row.Cells[2].String(),
					Address:        row.Cells[3].String(),
					VillageID:      toInt(row.Cells[4]),
					Phone:          row.Cells[5].String(),
					IsDeleted:      toInt(row.Cells[6]),
					IsActive:       toInt(row.Cells[7]),
					IsVerified:     toInt(row.Cells[8]),
					IsPoskora:      toInt(row.Cells[9]),
					IsDefault:      toInt(row.Cells[10]),
					SaldoBalance:   toFloat(row.Cells[11]),
					TotalOrder:     toInt(row.Cells[12]),
					CreatedBy:      toInt(row.Cells[13]),
					UpdatedBy:      toInt(row.Cells[14]),
					SurveyAnswered: toInt(row.Cells[15]),
				})
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println("[kk]success")

			errorCommit := tx.Commit()
			if errorCommit != nil {
				log.Fatalln(errorCommit)
			}
		}
		continue
	}
}

func InsertIntoUser(rows []*xlsx.Row, wg *sync.WaitGroup, db *sqlx.DB) {
	defer wg.Done()
	for i, row := range rows {
		if i > 0 {
			id, _ := strconv.Atoi(row.Cells[0].String())
			tx := db.MustBegin()
			query := "INSERT INTO user(id,usergroup_id,username,password,full_name,is_active,created_by,email,phone,is_superuser,is_verified,is_sales,android_id,saldo,catalog_profit,is_fnb)VALUES(:id,:usergroup_id,:username,:password,:full_name,:is_active,:created_by,:email,:phone,:is_superuser,:is_verified,:is_sales,:android_id,:saldo,:catalog_profit,:is_fnb)"
			_, err := tx.NamedExec(query,
				&model.User{
					ID:            id,
					UsergroupID:   toInt(row.Cells[1]),
					Username:      row.Cells[2].String(),
					Password:      row.Cells[3].String(),
					FullName:      row.Cells[4].String(),
					IsActive:      toInt(row.Cells[5]),
					CreatedBy:     toInt(row.Cells[6]),
					Email:         row.Cells[7].String(),
					Phone:         row.Cells[8].String(),
					IsSuperuser:   toInt(row.Cells[9]),
					IsVerified:    toInt(row.Cells[10]),
					IsSales:       toInt(row.Cells[11]),
					AndroidID:     row.Cells[12].String(),
					Saldo:         toFloat(row.Cells[13]),
					CatalogProfit: toFloat(row.Cells[14]),
					IsFnb:         toInt(row.Cells[15]),
				})
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println("[user]success")

			errorCommit := tx.Commit()
			if errorCommit != nil {
				log.Fatalln(errorCommit)
			}
		}
		continue
	}
}

func toInt(data *xlsx.Cell) int {
	result, _ := strconv.Atoi(data.String())
	return result
}

func toFloat(data *xlsx.Cell) float64 {
	result, _ := strconv.ParseFloat(data.String(), 64)
	return result
}

//Maret2020
