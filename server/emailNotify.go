/*
@Time : 2024/12/15 13:02
@Author : ljn
@File : emailNotify
@Software: GoLand
*/

package server

// EmailNotify 邮件通知
type EmailNotify struct {
	Username string
	Password string
	Host     string
	Port     int
	Receiver []string
}

// NewEmailNotify create a new email notify, and return the pointer
func NewEmailNotify(username, password, host string, port int, receiver []string) *EmailNotify {
	return &EmailNotify{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
		Receiver: receiver,
	}
}

func (e *EmailNotify) Send(title, content string) error {

	return nil
}
