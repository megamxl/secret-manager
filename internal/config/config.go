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
	Agent   AgentConfig   `mapstructure:"agent"`
	Server  ServerConfig  `mapstructure:"server"`
	Logging LoggingConfig `mapstructure:"logging"`
	DB      DBConfig      `mapstructure:"database"`
}

type AgentConfig struct {
	NodeName     string   `mapstructure:"nodeName"`
	WorkDir      string   `mapstructure:"workDir"`
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

func Load(configPath string) (*Config, error) {
	v := viper.New()

	// Defaults
	v.SetDefault("agent.workDir", "/var/lib/secret-manager")
	v.SetDefault("server.port", 8090)
	v.SetDefault("logging.level", "info")
	v.SetDefault("logging.format", "json")
	v.SetDefault("logging.auditFile", "/var/log/secret-manager/audit.log")
	v.SetDefault("database.path", "./secrets.db")

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
