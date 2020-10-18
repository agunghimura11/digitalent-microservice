package main

import (
	"fmt"
	"github.com/spf13/viper"
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

	//cfg := config.Config{
	//	Database : config.Database{
	//		Driver   : "mysql",
	//		Host     : "localhost",
	//		Port     : "3306",
	//		User     : "root",
	//		Password : "",
	//		DbName   : "digitalent_microservice",
	//		Config   : "charset=utf8&parseTime=True&loc=Local",
	//	},
	//}
	cfg, err := getConfig(); if err != nil {
		log.Println("Error get yaml", err.Error())
	}else{
		log.Println(cfg)
	}

	db, err := initDB(cfg.Database)
	if err != nil {
		log.Panic(err.Error())
	}else {
		log.Println("DB COnnection success")
	}

	router := mux.NewRouter()

	authMiddleware := handler.AuthMiddleware{
		AuthService: cfg.AuthService,
	}

	menuHandler := handler.Menu{Db: db}

	router.Handle("/menu", http.HandlerFunc(menuHandler.GetAllMenu))
	router.Handle("/add-menu", authMiddleware.ValidateAuth(http.HandlerFunc(menuHandler.AddMenu)))

	fmt.Printf("Server listen on :%s", cfg.Port)
	log.Panic(http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), router))
}

func getConfig() (config.Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.SetConfigName("config.yml")

	if err := viper.ReadInConfig(); err != nil {
		return config.Config{}, err
	}

	var cfg config.Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil
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

