# 消息推送工具

![Go version](https://img.shields.io/badge/go-%3E%3Dv1.22-9cf)
![Release](https://img.shields.io/badge/release-1.1-green.svg)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## 引入库
``` bash
go get github.com/ArtLjn/Notification@v0.1.0
```
## 钉钉推送使用
```go
func TestSendDingDing(t *testing.T) {
	// Example usage
	dingTalkSender := server.NewDingTalkSender("6f5321dc96471e1b396598fba6fd4133432a7778828c7fff8b881ad757b1f0683c9", "SECc8b27fb6ca615b650a505af14c408c312d833e83f9a27ac09161bbf34598fa46")
	context := util.NewNotificationContext(dingTalkSender)

	err := context.SendNotification("Test Title", "Test Content")
	if err != nil {
		fmt.Printf("Failed to send notification: %s\n", err)
	}
}
```
