package main

import (
	"digitalent-microservice/auth-service/config"
	"digitalent-microservice/auth-service/database"
	"fmt"
	"github.com/gorilla/mux"
	"digitalent-microservice/auth-service/handler"
	"log"
	"net/http"
    "github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

)

func main() {
	cfg, err := getConfig(); if err != nil {
		log.Println("Error get yaml",err.Error())
	}else{
		log.Println(cfg)
	}

	db,err := initDB(cfg.Database); if err != nil {
		log.Println(err.Error())
	}else{
		log.Println("DB COnnection success")
	}
	router := mux.NewRouter()
	authHandler := handler.AuthDB{Db: db}

	router.Handle("/auth/validate", http.HandlerFunc(authHandler.ValidateAuth))
	router.Handle("/auth/signup", http.HandlerFunc(authHandler.SignUp))
	router.Handle("/auth/login", http.HandlerFunc(authHandler.Login))

	fmt.Printf("Auth service listen on :8001")
	log.Panic(http.ListenAndServe(":8001", router))
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
	log.Println("Config",dbConfig)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName, dbConfig.Config)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&database.Auth{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
