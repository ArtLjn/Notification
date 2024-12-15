/*
@Time : 2024/12/14 22:01
@Author : ljn
@File : notification
@Software: GoLand
*/

package notify

// NotificationSender is the strategy interface for sending notifications.
type NotificationSender interface {
	Send(title, content string) error
}
