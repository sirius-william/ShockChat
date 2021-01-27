// @Title  InitConfig.go
// @Description	用来打印一些版权、logo的东西，注释掉了框架中Server.go中的printLogo方法
// @Author  SiriusWilliam  2021-01-20 10:58
// @Update  2021-01-20 10:58
/*
修改：
*/
package conf

import (
	"ShockChatServer/utils"
	"ShockChatServer/utils/mysql"
	uredis "ShockChatServer/utils/redis"
	"encoding/json"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
)

// Logo信息
var Logo string = `                                        
              ██                        
              ▀▀                        
 ████████   ████     ██▄████▄  ▀██  ██▀ 
     ▄█▀      ██     ██▀   ██    ████   
   ▄█▀        ██     ██    ██    ▄██▄   
 ▄██▄▄▄▄▄  ▄▄▄██▄▄▄  ██    ██   ▄█▀▀█▄  
 ▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀  ▀▀    ▀▀  ▀▀▀  ▀▀▀ 
                                        `

// 版权信息
var Copyright string = "[Github] https://github.com/aceld"

type PrintConfig struct {
	ShowLogo      bool `json:"show.logo"`
	ShowCopyright bool `json:"show.copyright"`
}

// 存储关于logo和版权的配置信息
var PrintConf PrintConfig = PrintConfig{true, true}

func init() {
	// 获取当前运行目录
	pwd, err := os.Getwd()
	// 检查logo.txt文件是否存在
	_, err2 := os.Stat(pwd + "/conf/logo.txt")
	// 如果存在，则读文件，替换logo
	if err2 == nil {
		file, _ := ioutil.ReadFile(pwd + "/conf/logo.txt")
		Logo = string(file)
	}
	// 检查copyright.txt文件是否存在
	_, err = os.Stat("./conf/copyright.txt")
	// 如果存在，则度=读整个文件
	if err == nil {
		file, _ := ioutil.ReadFile(pwd + "/conf/copyright.txt")
		Copyright = string(file)
	}
	// 读配置文件
	conf, err := ioutil.ReadFile(pwd + "/conf/zinx.json")
	// 反序列化到结构体变量
	if err == nil {
		_ = json.Unmarshal(conf, &PrintConf)
		_ = json.Unmarshal(conf, &mysql.MysqlConf)
		_ = json.Unmarshal(conf, &utils.KeyFile)
		_ = json.Unmarshal(conf, &utils.TokenConf)
		_ = json.Unmarshal(conf, &utils.EmailConf)
		_ = json.Unmarshal(conf, &uredis.RedisConf)
		// 初始化redis连接池
		uredis.RedisPool = uredis.InitRedisPool()
		utils.InitVars()
		mysql.InitMySqlPool()
	}
	// 显示
	if PrintConf.ShowLogo {
		color.Blue(Logo)
	}
	if PrintConf.ShowCopyright {
		color.HiBlue(Copyright)
	}
	if err == nil {
		color.Green("finding config file:%s, Unmarshall successfully.", pwd+"/conf/zinx.json")
	} else {
		color.Green("Didn't found config file, using default config.")
	}
}
