package zcsreceiver

import "fmt"

// Config represents the receiver config settings within the collector's config.yaml
type Config struct {
	Interval string   `mapstructure:"interval"`
	Auth     Auth     `mapstructure:"auth"`
	ThingIds []string `mapstructure:"thing_ids"`
}

// Auth represents the authentication settings for the ZCS receiver
type Auth struct {
	ClientID string `mapstructure:"client_id"`
	AuthKey  string `mapstructure:"auth_key"`
}

// Validate checks if the configuration is valid
func (cfg *Config) Validate() error {
	if cfg.Interval == "" {
		return fmt.Errorf("interval is required")
	}
	if cfg.Auth.ClientID == "" {
		return fmt.Errorf("client_id is required")
	}
	if cfg.Auth.AuthKey == "" {
		return fmt.Errorf("auth_key is required")
	}
	if len(cfg.ThingIds) == 0 {
		return fmt.Errorf("thing_ids is required")
	}

	return nil
}
