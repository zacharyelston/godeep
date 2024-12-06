// internal/client/client.go
package client

import (
	"fmt"
	"godeep/config"
	"net/http"
	"os"
	"time"
)

type DeepLakeClient struct {
	config *config.Config
	token  string
	client *http.Client
}

func NewDeepLakeClient(cfg *config.Config) (*DeepLakeClient, error) {
	token := os.Getenv("ACTIVELOOP_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("ACTIVELOOP_TOKEN environment variable not set")
	}

	transport := &http.Transport{
		MaxIdleConns:          cfg.Client.Transport.MaxIdleConns,
		IdleConnTimeout:       cfg.Client.Transport.IdleConnTimeout * time.Second,
		DisableCompression:    cfg.Client.Transport.DisableCompression,
		DisableKeepAlives:     cfg.Client.Transport.DisableKeepAlives,
		TLSHandshakeTimeout:   cfg.Client.Transport.TLSHandshakeTimeout * time.Second,
		ResponseHeaderTimeout: cfg.Client.Transport.ResponseHeaderTimeout * time.Second,
		ExpectContinueTimeout: cfg.Client.Transport.ExpectContinueTimeout * time.Second,
		ForceAttemptHTTP2:     cfg.Client.Transport.ForceHTTP2,
	}

	client := &http.Client{
		Timeout:   time.Duration(cfg.Client.Timeout) * time.Second,
		Transport: transport,
	}

	return &DeepLakeClient{
		config: cfg,
		token:  token,
		client: client,
	}, nil
}
