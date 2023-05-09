package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/celmonggo-api/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetConnection(ctx *gin.Context, config utils.Config) (db *mongo.Client) {
	mongoDSN := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.MongoDBUser, config.MongoDBPass, config.MongoDBHost, config.MongoDBPort)
	mongoconn := options.Client().ApplyURI(mongoDSN)
	db, err := mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	_ctx := context.Background()
	if err = db.Ping(_ctx, nil); err != nil {
		log.Fatalf("cannot connect to: %s [%v]", mongoDSN, err)
		return nil
	}

	_ctx.Done()
	return db
}

func GetConnectionTest() (db *sql.DB, error error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "Serverh5n&*#"
	dbName := "celestialdev"
	// datasource := "root:@tcp(localhost:3306)/celestialdev?parseTime=true"
	// db, err := sql.Open("mysql",dataSource)
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db, err
}

func GetConnectionDev() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "Serverh5n&*#"
	dbName := "celestialdev"
	// datasource := "root:@tcp(localhost:3306)/celestialdev?parseTime=true"
	// db, err := sql.Open("mysql",dataSource)
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}
