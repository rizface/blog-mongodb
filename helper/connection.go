package helper

import (
	"blog-mongo/app/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func Connection() (*mongo.Database,error){
	var err error
	var defaultConfig map[string]string
	var uri string
 	err = loadConfig()
	 if err != nil {
		 defaultConfig = config.DefaultConfig
		 uri = fmt.Sprintf("mongodb://%s:%s@%s:%s",defaultConfig["mongoUsername"],defaultConfig["mongoPassword"],defaultConfig["mongoHost"],defaultConfig["mongoPort"])
	 } else {
		 uri = fmt.Sprintf("mongodb://%s:%s@%s:%s",os.Getenv("MONGO_USERNAME"),os.Getenv("MONGO_PASSWORD"),os.Getenv("MONGO_HOST"),os.Getenv("MONGO_PORT"))
	 }
	 opts := options.Client().ApplyURI(uri).SetMaxPoolSize(50).SetMinPoolSize(10).SetMaxConnIdleTime(1 * time.Minute)
	 conn,err := mongo.Connect(context.TODO(),opts)
	 if err != nil {
		 log.Fatal(err.Error())
		 return nil,err
	 }
	 db := conn.Database("blog")
	 return db,nil
}