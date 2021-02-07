// @Title  MysqlInit.go
// @Description	数据库操作
// @Author  SiriusWilliam  2021-01-22 20:38
// @Update  2021-01-22 20:38
package mysql

import (
	//"database/sql"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type MySQLConfig struct {
	Host        string `json:"mysql.host"`
	UserName    string `json:"mysql.username"`
	Password    string `json:"mysql.password"`
	Port        int    `json:"mysql.port"`
	Database    string `json:"mysql.database"`
	Charset     string `json:"mysql.charset"`
	MaxConnSize int    `json:"mysql.maxConnSize"`
	MaxIdleSize int    `json:"mysql.maxIdleSize"`
}

var MysqlConf MySQLConfig = MySQLConfig{
	Host:        "127.0.0.1",
	Password:    "",
	Port:        3306,
	Database:    "shockchat",
	Charset:     "utf8mb4",
	MaxConnSize: 10,
	MaxIdleSize: 2,
}

var Mysql *sql.DB
var err error

func InitMySqlPool() {
	Host := MysqlConf.Host
	Password := MysqlConf.Password
	Port := MysqlConf.Port
	Database := MysqlConf.Database
	UserName := MysqlConf.UserName
	url := UserName + ":" + Password + "@tcp(" + Host + ":" + strconv.Itoa(Port) + ")/" + Database
	Mysql, err = sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	Mysql.SetMaxOpenConns(MysqlConf.MaxConnSize)
	Mysql.SetMaxIdleConns(MysqlConf.MaxIdleSize)
}
