package database

import (
	"context"
	"fmt"
	"log"

	"github.com/TalisonK/TalisonContabil/src/config"
	"github.com/TalisonK/TalisonContabil/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	Db *gorm.DB
}

var DBlocal DBInstance
var DBCloud *mongo.Client

func OpenConnectionLocal() error {

	conf := config.GetLocalDB()

	fmt.Println(conf)

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

	conn.AutoMigrate(&model.Category{})
	conn.AutoMigrate(&model.User{})
	conn.AutoMigrate(&model.Income{})
	conn.AutoMigrate(&model.Expense{})
	conn.AutoMigrate(&model.List{})

	DBlocal = DBInstance{
		Db: conn,
	}

	return nil
}

func OpenConnectionCloud() error {

	conf := config.GetCloudDB()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	dsn := fmt.Sprintf("%s://%s:%s@%s/?retryWrites=true&w=majority&appName=Base-contabil", conf.Host, conf.User, conf.Pass, conf.Database)

	opts := options.Client().ApplyURI(dsn).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		log.Fatalln("Failed to connect to cloud database.", err)
		return err
	}

	// enviando um ping para confirmar conexão
	err = client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err()

	test := client.Database("contabil").Collection("category")

	curs, err := test.Find(context.TODO(), bson.D{})

	for curs.Next(context.TODO()) {
		var result bson.M
		curs.Decode(&result)
		fmt.Println(result)
	}

	if err != nil {
		log.Fatalln("Failed to ping cloud database.", err)
		return err
	}

	DBCloud = client

	return nil
}

func CloseConnections() {
	db, _ := DBlocal.Db.DB()

	err := db.Close()

	if err != nil {
		log.Fatalln("Failed to close local database.", err)
	}
}
