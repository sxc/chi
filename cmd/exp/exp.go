package main

import (
	"fmt"

	"github.com/sxc/oishifood/models"
)

func DefaultPostgresConfig() PostgresConfig {

	return PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "oishifooduser",
		Password: "oishifoodpassword",
		Database: "oishifooddb",
		SSLMode:  "disable",
	}
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

func main() {
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")

	us := models.UserService{
		DB: db,
	}
	user, err := us.Create("abcd2@abcd.com", "abc123")
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
