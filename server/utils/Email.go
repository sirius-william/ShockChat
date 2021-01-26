// @Title  Email
// @Description	邮箱发邮件模块
// @Author  SiriusWilliam  2021-01-26 0:08
// @Update  2021-01-26 0:08
package utils

import (
	"errors"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"os"
	"strings"
)

type EmailConfig struct {
	Name         string            `json:"mail.name"`
	MailSmtpHost string            `json:"mail.host"`     // 邮箱服务器地址
	MailUser     string            `json:"mail.user"`     // 邮箱地址
	MailPassword string            `json:"mail.password"` // 邮箱密码
	Template     map[string]string `json:"mail.template"` // 模板，必须存放在template文件夹内
}

var EmailConf EmailConfig

func init() {
	EmailConf = EmailConfig{"ShockChat", "", "", "", make(map[string]string)}
}

/*
 * 通过templateId来读取模板，用mapper里的内容来替换模板中的保留字
 */
func SendMail(templateId string, mapper map[string]string, toUser string, title string) error {
	var templateString string
	if len(EmailConf.Template) == 0 {
		templateString = "你的验证码为{{ number }}，有效时间5分钟。"
	} else {
		// 获取绝对地址
		pwd, _ := os.Getwd()
		filePath := pwd + "/template/" + EmailConf.Template[templateId]
		// 读到指定模板文件
		templateFile, err := ioutil.ReadFile(filePath)
		if err != nil {
			return err
		}
		// 转成字符串
		templateString = string(templateFile)
	}
	// 遍历mapper，将mapper中的字段替换到模板中
	for contain := range mapper {
		templateString = strings.Replace(string(templateString), "{{ "+contain+" }}", mapper[contain], -1)
	}
	// 没有内容，则返回
	if templateString == "" {
		return errors.New("data need to send is empty")
	}
	// 开始发邮件
	mail := gomail.NewMessage()
	mail.SetHeader("From", EmailConf.Name+"<"+EmailConf.MailUser+">") // 发件人
	mail.SetHeader("To", toUser)                                      // 收件人
	mail.SetHeader("Subject", title)                                  // 标题
	mail.SetBody("text/html", templateString)                         // 内容
	// 打包
	send := gomail.Dialer{
		Host:     EmailConf.MailSmtpHost,
		Port:     25,
		Username: EmailConf.MailUser,
		Password: EmailConf.MailPassword,
		SSL:      false,
	}
	// 发送
	if err := send.DialAndSend(mail); err != nil {
		return err
	}
	return nil
}
