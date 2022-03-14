package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kevin19930919/CryptoAlert/model"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     7552,
		User:     "root",
		Password: "kevin7552",
		DBName:   "cryptoalert",
	}
	return &dbConfig
}

func InitDB(dbConfig *DBConfig) string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.DBName,
		dbConfig.Password,
	)
}

func StartPostgrel() error {
	constr := InitDB(BuildDBConfig())
	DB, err := gorm.Open("postgres", constr)
	if err != nil {
		fmt.Println("fail to connect database:%s", err)
	}

	DB.AutoMigrate(&model.Alert{})

	return err
}
