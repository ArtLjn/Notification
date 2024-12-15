/*
@Time : 2024/12/15 13:02
@Author : ljn
@File : emailNotify
@Software: GoLand
*/

package server

import "gopkg.in/gomail.v2"

// EmailNotify 邮件通知
type EmailNotify struct {
	Username string
	Password string
	Host     string
	Port     int
	Receiver []string
}

// NewEmailNotify 创建一个新的 EmailNotify 实例
func NewEmailNotify(username, password, host string, port int, receiver []string) *EmailNotify {
	return &EmailNotify{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
		Receiver: receiver,
	}
}

// Send 使用 SSL 发送邮件
func (e *EmailNotify) Send(title, content string) error {
	// 创建邮件
	m := gomail.NewMessage()
	m.SetHeader("From", e.Username) // 确保发件人邮箱格式正确
	m.SetHeader("To", e.Receiver...)
	m.SetHeader("Subject", title)
	m.SetBody("text/plain", content)

	// 创建 SMTP 连接
	d := gomail.NewDialer(e.Host, e.Port, e.Username, e.Password)

	// 发送邮件
	return d.DialAndSend(m)
}
