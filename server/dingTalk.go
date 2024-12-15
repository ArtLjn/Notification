/*
@Time : 2024/12/14 22:01
@Author : ljn
@File : DingTalk
@Software: GoLand
*/

package server

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/ArtLjn/Notification/util"
	"net/http"
	"net/url"
	"time"
)

// DingTalkSender is a concrete implementation of NotificationSender for DingTalk.
type DingTalkSender struct {
	Token  string
	Secret string
}

// NewDingTalkSender creates a new instance of DingTalkSender.
func NewDingTalkSender(token, secret string) *DingTalkSender {
	return &DingTalkSender{
		Token:  token,
		Secret: secret,
	}
}

// Send sends a notification using DingTalk.
func (d *DingTalkSender) Send(title, content string) error {
	if d.Token == "" || d.Secret == "" {
		return fmt.Errorf("DingTalk token or secret is not set")
	}

	timestamp := time.Now().UnixMilli()
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, d.Secret)

	hmacHashed := hmac.New(sha256.New, []byte(d.Secret))
	hmacHashed.Write([]byte(stringToSign))
	sign := url.QueryEscape(base64.StdEncoding.EncodeToString(hmacHashed.Sum(nil)))

	apiURL := fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s&timestamp=%d&sign=%s", d.Token, timestamp, sign)

	data := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": fmt.Sprintf("%s\n\n%s", title, content),
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal request data: %w", err)
	}

	headers := map[string]string{
		"Content-Type": "application/json;charset=utf-8",
	}

	response, err := util.MakeHTTPRequest(http.MethodPost, apiURL, jsonData, headers)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	var respBody map[string]interface{}
	if err := json.Unmarshal(response, &respBody); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if errCode, ok := respBody["errcode"].(float64); ok && errCode != 0 {
		return fmt.Errorf("DingTalk notification failed with errcode: %.0f", errCode)
	}

	fmt.Println("DingTalk notification sent successfully!")
	return nil
}
