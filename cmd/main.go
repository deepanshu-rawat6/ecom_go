package main

import (
	"log"

	"github.com/deepanshu-rawat6/ecom_go/cmd/api"
	"github.com/deepanshu-rawat6/ecom_go/config"
	"github.com/deepanshu-rawat6/ecom_go/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
