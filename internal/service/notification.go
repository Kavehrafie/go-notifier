package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type NotificationService interface {
	Send(ctx context.Context, url string, payload interface{}) (int, error)
}

type httpNotificationService struct {
	client *http.Client
}

func (s *httpNotificationService) Send(ctx context.Context, url string, payload interface{}) (int, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal payload: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return 0, fmt.Errorf("failed to create http request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to send http request: %v", err)
	}
	defer resp.Body.Close()

	// Read and discard body to reuse connection
	_, _ = io.Copy(io.Discard, resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return resp.StatusCode, fmt.Errorf("webhook returned status code %d", resp.StatusCode)
	}

	return resp.StatusCode, nil
}

func NewNotificationService() NotificationService {
	return &httpNotificationService{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
