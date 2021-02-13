// @Title  Redis.go
// @Description
// @Author  SiriusWilliam  2021-01-26 11:27
// @Update  2021-01-26 11:27
package redis

import (
	"ShockChatServer/logger"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
)

type RedisConfig struct {
	Host              string `json:"redis.host"`
	Password          string `json:"redis.password"`
	Port              int    `json:"redis.port"`
	MaxIdle           int    `json:"redis.max_idle"`
	MaxConnSize       int    `json:"redis.max_conn_size"`
	ReadTimeout       int    `json:"redis.read_timeout"`
	WriteTimeout      int    `json:"redis.write_timeout"`
	Database          int    `json:"redis.database"`
	ConnectionTimeout int    `json:"redis.connection_timeout"`
}

var RedisPool *redis.Pool

var RedisConf RedisConfig = RedisConfig{
	Host:              "127.0.0.1",
	Password:          "",
	Port:              3306,
	MaxIdle:           3,
	MaxConnSize:       100,
	ReadTimeout:       10,
	WriteTimeout:      10,
	Database:          0,
	ConnectionTimeout: 10,
}

func InitRedisPool() {
	logger.Log.INFO("Init RedisPool: Connect to " + RedisConf.Host + ":" + strconv.Itoa(RedisConf.Port))
	RedisPool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			dial, err := redis.Dial(
				"tcp",
				RedisConf.Host+":"+strconv.Itoa(RedisConf.Port),
				redis.DialReadTimeout(time.Duration(RedisConf.ReadTimeout)*time.Second),
				redis.DialWriteTimeout(time.Duration(RedisConf.WriteTimeout)*time.Second),
				redis.DialConnectTimeout(time.Duration(RedisConf.ConnectionTimeout)*time.Second),
				redis.DialDatabase(RedisConf.Database),
			)
			if err != nil {
				panic(err)
			}
			if RedisConf.Password != "" {
				_, err := dial.Do("auth", RedisConf.Password)
				if err != nil {
					panic(err)
				}
			}
			return dial, nil
		},
		DialContext:     nil,
		TestOnBorrow:    nil,
		MaxIdle:         0,
		MaxActive:       0,
		IdleTimeout:     0,
		Wait:            false,
		MaxConnLifetime: 0,
	}
}

func Exec(cmd string, args ...interface{}) ([]string, error) {
	client := RedisPool.Get()
	defer client.Close()
	res, err := client.Do(cmd, args...)
	if err != nil {
		return nil, err
	}
	var resList []string = make([]string, 1)
	if value, ok := res.([]byte); ok {
		resList = append(resList, string(value))
	}
	if value, ok := res.([]interface{}); ok {
		for _, v := range value {
			resList = append(resList, string(v.([]byte)))
		}
	}
	return resList, nil
}

func DisConnect() error {
	err := RedisPool.Close()
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}
	return nil
}
