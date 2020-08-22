package config

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("config_example")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
}

func DBConnect() *sqlx.DB {
	DB_USER := viper.GetString("db.user")
	DB_PASSWORD := viper.GetString("db.password")
	DB_HOST := viper.GetString("db.host")
	DB_PORT := viper.GetString("db.port")
	DB_NAME := viper.GetString("db.name")

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	db, err := sqlx.Connect("mysql", connection)
	if err != nil {
		log.Fatalln(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	return db
}
