// @Title  MongoDBInit.go
// @Description
// @Author  SiriusWilliam  2021-02-03 11:03
// @Update  2021-02-03 11:03
package mongodb

import (
	"ShockChatServer/logger"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"strconv"
	"time"
)

type MongoConfiguration struct {
	Host          string `json:"mongodb.host"`
	UserName      string `json:"mongodb.username"`
	Password      string `json:"mongodb.password"`
	Database      string `json:"mongodb.database"`
	Port          int    `json:"mongodb.port"`
	TimeOut       int    `json:"mongodb.timeout"`
	AuthMechanism string `json:"mongodb.authMechanism"`
}

var authMechanism []string = []string{"MONGODB-AWS"}

var MongoConfig = &MongoConfiguration{
	Host:          "localhost",
	UserName:      "root",
	Password:      "",
	Database:      "test",
	Port:          27017,
	TimeOut:       10,
	AuthMechanism: "MONGODB-AWS",
}
var MongoCtx context.Context
var cancel context.CancelFunc = func() {
	logger.Log.INFO("cancel")
}
var client *mongo.Client

func InitMongoDB() {
	var err error
	var url string = "mongodb://" + MongoConfig.Host + ":" + strconv.Itoa(MongoConfig.Port)
	logger.Log.INFO("Init MongoDB: Connect to " + url)
	MongoCtx, cancel = context.WithTimeout(context.Background(), time.Duration(MongoConfig.TimeOut)*time.Second)
	defer cancel()
	var opt *options.ClientOptions = options.Client().ApplyURI(url)
	var isAuthMechanismContains bool
	for _, auth := range authMechanism {
		if auth == MongoConfig.AuthMechanism {
			isAuthMechanismContains = true
			break
		}
	}
	if isAuthMechanismContains && MongoConfig.Password != "" {
		var credential options.Credential = options.Credential{
			AuthMechanism: MongoConfig.AuthMechanism,
			Password:      MongoConfig.Password,
			Username:      MongoConfig.UserName,
		}
		opt.SetAuth(credential)
	}
	client, err = mongo.Connect(MongoCtx, opt)
	if err != nil {
		panic(err)
	}
	err = client.Ping(MongoCtx, readpref.Primary())
	if err != nil {
		panic(err)
	}
}

func DisConnect() error {
	err := client.Disconnect(MongoCtx)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}
	return nil
}
