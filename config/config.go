// config/config.go
package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Activeloop ActiveloopConfig `mapstructure:"activeloop"`
	Client     ClientConfig     `mapstructure:"client"`
}

type ActiveloopConfig struct {
	BaseURL       string                 `mapstructure:"base_url"`
	OrgID         string                 `mapstructure:"org_id"`
	DatasetPath   string                 `mapstructure:"dataset_path"`
	DefaultSchema map[string]interface{} `mapstructure:"default_schema"`
}

type ClientConfig struct {
	Timeout    int             `mapstructure:"timeout"`
	MaxRetries int             `mapstructure:"max_retries"`
	RetryDelay time.Duration   `mapstructure:"retry_delay"`
	Transport  TransportConfig `mapstructure:"transport"`
}

type TransportConfig struct {
	MaxIdleConns          int           `mapstructure:"max_idle_conns"`
	IdleConnTimeout       time.Duration `mapstructure:"idle_conn_timeout"`
	DisableCompression    bool          `mapstructure:"disable_compression"`
	DisableKeepAlives     bool          `mapstructure:"disable_keep_alives"`
	TLSHandshakeTimeout   time.Duration `mapstructure:"tls_handshake_timeout"`
	ResponseHeaderTimeout time.Duration `mapstructure:"response_header_timeout"`
	ExpectContinueTimeout time.Duration `mapstructure:"expect_continue_timeout"`
	ForceHTTP2            bool          `mapstructure:"force_http2"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("default")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("ACTIVELOOP")
	viper.AutomaticEnv()

	viper.BindEnv("activeloop.token", "ACTIVELOOP_TOKEN")
	viper.BindEnv("activeloop.org_id", "ACTIVELOOP_ORG_ID")
	viper.BindEnv("activeloop.dataset_path", "ACTIVELOOP_DATASET_PATH")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %v", err)
	}

	return &config, nil
}
