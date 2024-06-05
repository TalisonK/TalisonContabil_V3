package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/TalisonK/TalisonContabil/src/config"
	"github.com/TalisonK/TalisonContabil/src/domain"
	"github.com/TalisonK/TalisonContabil/src/util/constants"
	"github.com/TalisonK/TalisonContabil/src/util/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type CloudCollections struct {
	Base     *mongo.Client
	Category *mongo.Collection
	Expense  *mongo.Collection
	Income   *mongo.Collection
	List     *mongo.Collection
	User     *mongo.Collection
	Total    *mongo.Collection
}

var DBlocal *gorm.DB
var DBCloud CloudCollections

// OpenConnectionLocal starts a connection with the local database
func OpenConnectionLocal() error {

	conf := config.GetLocalDB()

	fmt.Println(conf)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Pass, conf.Host, conf.Port, conf.Database)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,  // Slow SQL threshold
			LogLevel:      logger.Error, // Log level
			Colorful:      false,        // Disable color
		},
	)

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})

	if err != nil {
		return fmt.Errorf(logging.FailedToOpenConnection(constants.LOCAL, err))
	}

	logging.OpenedConnection(constants.LOCAL)
	conn.Logger = logger.Default.LogMode(logger.Info)

	conn.AutoMigrate(&domain.Category{})
	conn.AutoMigrate(&domain.User{})
	conn.AutoMigrate(&domain.Income{})
	conn.AutoMigrate(&domain.Expense{})
	conn.AutoMigrate(&domain.List{})
	conn.AutoMigrate(&domain.Total{})

	DBlocal = conn

	return nil
}

// OpenConnectionCloud starts a connection with the cloud database
func OpenConnectionCloud() error {

	conf := config.GetCloudDB()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	dsn := fmt.Sprintf("%s://%s:%s@%s/?retryWrites=true&w=majority&appName=Base-contabil", conf.Host, conf.User, conf.Pass, conf.Database)

	opts := options.Client().ApplyURI(dsn).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		logging.FailedToOpenConnection(constants.CLOUD, err)
		return err
	}

	// enviando um ping para confirmar conexão
	err = client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err()

	if err != nil {
		logging.FailedToConnectToDB(constants.CLOUD, err)
		return err
	}

	var base string

	if config.IsProd() {
		base = "contabil"
	} else {
		base = "contabil-dev"
	}

	DBCloud.Base = client
	DBCloud.User = client.Database(base).Collection("user")
	DBCloud.Category = client.Database(base).Collection("category")
	DBCloud.Income = client.Database(base).Collection(constants.INCOME)
	DBCloud.Expense = client.Database(base).Collection(constants.EXPENSE)
	DBCloud.List = client.Database(base).Collection("list")
	DBCloud.Total = client.Database(base).Collection("total")

	logging.OpenedConnection(constants.CLOUD)
	return nil
}

// CheckLocalDB checks if the local database is connected
func checkLocalDB() bool {

	if DBlocal == nil {
		logging.FailedToPingDB(constants.LOCAL, nil)
		return false
	}

	section, err := DBlocal.DB()

	if err != nil {
		logging.FailedToConnectToDB(constants.LOCAL, err)
		return false
	}

	err = section.Ping()

	if err != nil {
		logging.FailedToPingDB(constants.LOCAL, err)
		return false
	}
	return true

}

// CheckCloudDB checks if the cloud database is connected
func checkCloudDB() bool {

	err := DBCloud.Base.Ping(context.Background(), nil)

	if err != nil {
		logging.FailedToPingDB(constants.CLOUD, nil)
		return false
	}

	return true
}

// CloseConnections closes the connections with the databases
func CloseConnections() {
	db, _ := DBlocal.DB()

	err := db.Close()

	if err != nil {
		logging.FailedToCloseConnection(constants.LOCAL, err)
	}
}

func CheckDBStatus() (bool, bool) {
	// Check database status
	statusDbLocal := checkLocalDB()
	statusDbCloud := checkCloudDB()

	return statusDbLocal, statusDbCloud
}
