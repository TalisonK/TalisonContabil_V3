package database

import (
	"fmt"
	"log"

	"github.com/TalisonK/media-storager/src/config"
	"github.com/TalisonK/media-storager/src/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	Db *gorm.DB
}

var DB DBInstance

func OpenConnectionLocal() error {

	conf := config.GetLocalDB()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Pass, conf.Host, conf.Port, conf.Database)

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalln("Failed to connect to database.", err)
		return err
	}

	log.Println("Connected to database!")
	conn.Logger = logger.Default.LogMode(logger.Info)

	conn.AutoMigrate(&model.Media{})

	DB = DBInstance{
		Db: conn,
	}

	return nil
}

func CloseConnection() {
	db, err := DB.Db.DB()
	if err != nil {
		log.Fatalln("Failed to close database.", err)
	}
	db.Close()
}
