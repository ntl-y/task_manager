package main

import (
	"os"
	taskmanager "task_manager"
	"task_manager/pkg/handler"
	"task_manager/pkg/repository"
	"task_manager/pkg/service"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initConfig() error {
	if err := godotenv.Load("configs/.env"); err != nil {
		return err
	}

	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalln(err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalln(err)
	} else {
		logrus.Println("Connected to db")
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	server := taskmanager.NewServer(viper.GetString("port"), handler.InitRoutes())

	if err := server.Run(); err != nil {
		logrus.Fatalln(err)
	}

}
