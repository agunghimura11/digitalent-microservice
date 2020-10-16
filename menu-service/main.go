package main

import (
	"fmt"
	"log"
	"net/http"
	"digitalent-microservice/menu-service/config"
	"digitalent-microservice/menu-service/handler"
	"digitalent-microservice/menu-service/database"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

)

func main(){

	cfg := config.Config{
		Database : config.Database{
			Driver   : "mysql",
			Host     : "localhost",
			Port     : "3306",
			User     : "root",
			Password : "",
			DbName   : "digitalent_microservice",
			Config   : "charset=utf8&parseTime=True&loc=Local",
		},
	}

	db, err := initDB(cfg.Database)
	if err != nil {
		log.Panic(err)
		return
	}

	router := mux.NewRouter()

	authMiddleware := handler.AuthMiddleware{
		AuthService: cfg.AuthService,
	}

	menuHandler := handler.Menu{Db: db}

	router.Handle("/add-menu", http.HandlerFunc(menuHandler.AddMenu))
	router.Handle("/menu", authMiddleware.ValidateAuth(http.HandlerFunc(menuHandler.GetAllMenu)))

	fmt.Printf("Auth service listen on :8001")
	log.Panic(http.ListenAndServe(":8001", router))
}

func initDB(dbConfig config.Database) (*gorm.DB, error) {
	log.Println("config", dbConfig)
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName,dbConfig.Config)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName,dbConfig.Config)
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

