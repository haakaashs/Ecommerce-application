package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Conn   *gorm.DB
	Config Env
	err    error
)

// Environment struct
type Env struct {
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
}

// Load env from .env file
func loadEnv(path string) Env {
	err = godotenv.Load(path)
	if err != nil {
		log.Fatal("cannot load env: ", err)

	}

	return Env{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		ServerPort: os.Getenv("SERVER_PORT"),
	}
}

// Initialize database connection
func InitializeDB(path string) {

	Config = loadEnv(path)
	dns := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", Config.DBUser, Config.DBPassword, Config.DBName)
	Conn, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("unable to connect DB: ", err.Error())
	}

	if d, ok := Conn.DB(); ok != nil {
		if err = d.Ping(); err != nil {
			d.Close()
			log.Fatal(err)

		}
	}

	log.Println("DB Connected!")
}
