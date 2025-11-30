package main

import (
	"fmt"
	"log"
	"os"

	server "github.com/balatsanandrey25/PWG"
	"github.com/balatsanandrey25/PWG/pkg/controllers"
	"github.com/balatsanandrey25/PWG/pkg/handler"
	"github.com/balatsanandrey25/PWG/pkg/models"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := models.ConfigurationInstaller(); err != nil {
		log.Fatal(err)
	}
	db, err := models.NewPostgresDB(models.Config{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		Username: viper.GetString("database.username"),
		DBName:   viper.GetString("database.dbname"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		SSLMode:  viper.GetString("database.sslmode"),
	})
	if err != nil {
		log.Fatalf("error init DB %s", err.Error())
	}

	model := models.NewModels(db)
	controllres := controllers.NewControllers(model)
	handlers := handler.NewHandler(controllres)
	fmt.Println("The server is running on the port:", viper.GetString("startserver.port"))
	srv := new(server.Server)
	if err := srv.Run(viper.GetString("server.port"), handlers.InitHandler()); err != nil {
		log.Fatalf("The server fell to the death of the brave: %s", err)
	}
}
