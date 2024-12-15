/*
@Time : 2024/12/14 22:19
@Author : ljn
@File : send_notification_test
@Software: GoLand
*/

package _example

import (
	"fmt"
	"github.com/ArtLjn/Notification/server"
	"github.com/ArtLjn/Notification/util"
	"testing"
)

func TestSendDingDing(t *testing.T) {
	// Example usage
	dingTalkSender := server.NewDingTalkSender("6f5321dc96471e1b396598fba6fd4133432a7778828c7fff8b881ad757b1f0683c9", "SECc8b27fb6ca615b650a505af14c408c312d833e83f9a27ac09161bbf34598fa46")
	context := util.NewNotificationContext(dingTalkSender)

	err := context.SendNotification("Test Title", "Test Content")
	if err != nil {
		fmt.Printf("Failed to send notification: %s\n", err)
	}
}
