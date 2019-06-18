package config

// Config defines the configuration structure.
type MxpConfig struct {
	General struct {
		LogLevel               int    `mapstructure:"log_level"`
		PasswordHashIterations int    `mapstructure:"password_hash_iterations"`
		HostServer             string `mapstructure:"host_server"`
		BindPort               string `mapstructure:"bind_port"`
	}

	PostgreSQL struct {
		DSN         string `mapstructure:"dsn"`
		Automigrate bool   `mapstructure:"automigrate"`
	}
}

// C holds the global configuration.
var Cstruct MxpConfig
