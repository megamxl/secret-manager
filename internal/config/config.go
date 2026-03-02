package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

import "sync"

var (
	instance *Config
	once     sync.Once
)

type Config struct {
	Agent          AgentConfig    `mapstructure:"agent"`
	Server         ServerConfig   `mapstructure:"server"`
	Logging        LoggingConfig  `mapstructure:"logging"`
	DB             DBConfig       `mapstructure:"database"`
	SecurityConfig SecurityConfig `mapstructure:"security"`
}

type AgentConfig struct {
	NodeName     string   `mapstructure:"nodeName"`
	TrustedRoots []string `mapstructure:"trustedRoots"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type LoggingConfig struct {
	Level     string `mapstructure:"level"`     // debug, info, warn, error
	Format    string `mapstructure:"format"`    // json or console
	AuditFile string `mapstructure:"auditFile"` // path to audit log
}

type DBConfig struct {
	Path string `mapstructure:"path"`
}

type SecurityConfig struct {
	TLS  TLSConfig  `mapstructure:"tls"`
	Auth AuthConfig `mapstructure:"auth"`
}

type TLSConfig struct {
	Enabled  bool   `mapstructure:"enabled"`
	CertFile string `mapstructure:"certFile"`
	KeyFile  string `mapstructure:"keyFile"`
	// Used for mTLS client CA verification
	ClientCAFile string `mapstructure:"clientCAFile"`
}

type AuthConfig struct {
	// "none", "mtls", "jwt", "spire"
	Method string      `mapstructure:"method"`
	JWT    JWTConfig   `mapstructure:"jwt"`
	SPIRE  SPIREConfig `mapstructure:"spire"`
}

type JWTConfig struct {
	// Provide either a static secret (HMAC) or a JWKS URL (RSA/EC)
	Secret   string `mapstructure:"secret"`   // HMAC shared secret
	JWKSURL  string `mapstructure:"jwksUrl"`  // e.g. https://issuer/.well-known/jwks.json
	Issuer   string `mapstructure:"issuer"`   // validated if non-empty
	Audience string `mapstructure:"audience"` // validated if non-empty
}

type SPIREConfig struct {
	SocketPath  string `mapstructure:"socketPath"`  // /tmp/spire-agent/public/api.sock
	TrustDomain string `mapstructure:"trustDomain"` // e.g. example.org
	// Allowlist of SPIFFE IDs that may call the API; empty = any valid SVID
	AllowedSPIFFEIDs []string `mapstructure:"allowedSpiffeIds"`
}

func Load(configPath string) (*Config, error) {
	v := viper.New()

	// Defaults
	v.SetDefault("server.port", 8090)
	v.SetDefault("logging.level", "info")
	v.SetDefault("logging.format", "json")
	v.SetDefault("logging.auditFile", "/var/log/secret-manager/audit.log")
	v.SetDefault("database.path", "./secrets.db")
	v.SetDefault("security.tls.enabled", false)
	v.SetDefault("security.auth.method", "none")

	// Load file
	if configPath != "" {
		v.SetConfigFile(configPath)
		if err := v.ReadInConfig(); err != nil {
			return nil, fmt.Errorf("failed to read config: %w", err)
		}
	}

	// Environment variables override
	v.SetEnvPrefix("SECRET_MANAGER")
	v.AutomaticEnv()

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	// Auto-detect nodeName if not set
	if cfg.Agent.NodeName == "" {
		hostname, _ := os.Hostname()
		cfg.Agent.NodeName = hostname
	}

	return &cfg, nil
}

func Initialize(configPath string) error {
	var err error
	once.Do(func() {
		instance, err = Load(configPath)
	})
	return err
}

func Get() *Config {
	if instance == nil {
		panic("config not initialized - call Initialize() first")
	}
	return instance
}
