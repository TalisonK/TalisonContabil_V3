package database

import (
	"context"
	"fmt"

	"github.com/TalisonK/TalisonContabil/src/config"
	"github.com/TalisonK/TalisonContabil/src/model"
	"github.com/TalisonK/TalisonContabil/src/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type CloudCollections struct {
	User     *mongo.Collection
	Category *mongo.Collection
	Income   *mongo.Collection
	Expense  *mongo.Collection
	List     *mongo.Collection
}

var DBlocal *gorm.DB
var DBCloud CloudCollections

func OpenConnectionLocal() error {

	conf := config.GetLocalDB()

	fmt.Println(conf)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Pass, conf.Host, conf.Port, conf.Database)

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		util.LogHandler("Failed to connect to database.", err, "OpenConnectionLocal")
		return err
	}

	util.LogHandler("Connected to database!", nil, "OpenConnectionLocal")
	conn.Logger = logger.Default.LogMode(logger.Info)

	conn.AutoMigrate(&model.Category{})
	conn.AutoMigrate(&model.User{})
	conn.AutoMigrate(&model.Income{})
	conn.AutoMigrate(&model.Expense{})
	conn.AutoMigrate(&model.List{})

	DBlocal = conn

	return nil
}

func OpenConnectionCloud() error {

	conf := config.GetCloudDB()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	dsn := fmt.Sprintf("%s://%s:%s@%s/?retryWrites=true&w=majority&appName=Base-contabil", conf.Host, conf.User, conf.Pass, conf.Database)

	opts := options.Client().ApplyURI(dsn).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		util.LogHandler("Failed to connect to cloud database.", err, "OpenConnectionCloud")
		return err
	}

	// enviando um ping para confirmar conexão
	err = client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err()

	if err != nil {
		util.LogHandler("Failed to ping cloud database.", err, "OpenConnectionCloud")
		return err
	}

	DBCloud.User = client.Database("contabil").Collection("user")
	DBCloud.Category = client.Database("contabil").Collection("category")
	DBCloud.Income = client.Database("contabil").Collection("income")
	DBCloud.Expense = client.Database("contabil").Collection("expense")
	DBCloud.List = client.Database("contabil").Collection("list")

	return nil
}

func CheckLocalDB() bool {
	section, err := DBlocal.DB()

	if err != nil {
		util.LogHandler("Failed to connect to local database.", err, "checkLocalDB")
		return false
	}

	err = section.Ping()

	if err != nil {
		util.LogHandler("Failed to ping local database.", err, "checkLocalDB")
		return false
	}
	return true

}

func CheckCloudDB() bool {
	if DBCloud.Expense == nil {
		util.LogHandler("Failed to ping cloud database.", nil, "checkCloudDB")
		return false
	}

	err := DBCloud.Expense.FindOne(context.TODO(), bson.D{}).Err()

	if err != nil {
		util.LogHandler("Failed to ping cloud database.", err, "checkCloudDB")
		return false
	}
	return true
}

func CloseConnections() {
	db, _ := DBlocal.DB()

	err := db.Close()

	if err != nil {
		util.LogHandler("Failed to close local database connection.", err, "CloseConnections")
	}
}
