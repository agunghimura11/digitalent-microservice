package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"digitalent-microservice/menu-service/handler"
)

func main(){

	cfg := config.Config{
		Database : config.Database{
			Driver   : "mysql",
			Host     : "localhost",
			Port     : "3306",
			User     : "root",
			Password : "",
			DbName   : "test",
			Config   :`charset=utf8&parseTime=True&loc=Local`,
		},
	}
	
	db, err := initDB(cfg.Database)
	if err != nil {
		log.Panic(err)
		return
	} 

	router := mux.NewRouter()

	router.Handle("/add-menu", http.HandlerFunc(handler.AddMenu))

	fmt.Println("Menu service listen on port :8000")
	log.Panic(http.ListenAndServe(":8000", router))
}

func initDB(dbConfig config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Config)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(database.Menu{})
	if err != nil {
		return nil, err
	}
	return db, nil
}