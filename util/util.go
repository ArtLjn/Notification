/*
@Time : 2024/12/14 22:02
@Author : ljn
@File : util
@Software: GoLand
*/

package util

import (
	"bytes"
	"fmt"
	"github.com/ArtLjn/Notification"
	"io"
	"net/http"
	"time"
)

// NotificationContext is the context that uses a NotificationSender.
type NotificationContext struct {
	sender notify.NotificationSender
}

// NewNotificationContext creates a new NotificationContext.
func NewNotificationContext(sender notify.NotificationSender) *NotificationContext {
	return &NotificationContext{
		sender: sender,
	}
}

// SendNotification sends a notification using the configured sender.
func (c *NotificationContext) SendNotification(title, content string) error {
	return c.sender.Send(title, content)
}

// MakeHTTPRequest is a helper function to make HTTP requests.
func MakeHTTPRequest(method, url string, body []byte, headers map[string]string) ([]byte, error) {
	client := &http.Client{Timeout: 15 * time.Second}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return respBody, nil
}
