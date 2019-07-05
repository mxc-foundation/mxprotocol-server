package config

// Config defines the configuration structure.
type MxpConfig struct {
	General struct {
		LogLevel               int    `mapstructure:"log_level"`
		PasswordHashIterations int    `mapstructure:"password_hash_iterations"`
		HostServer             string `mapstructure:"host_server"`
		AuthServer             string `mapstructure:"auth_server"`
		AuthUrl                string `mapstructure:"auth_url"`
	}

	PostgreSQL struct {
		DSN         string `mapstructure:"dsn"`
		Automigrate bool   `mapstructure:"automigrate"`
	}

	ApplicationServer struct {
		HttpServer struct {
			Bind                       string `mapstructure:"bind"`
			TLSCert                    string `mapstructure:"tls_cert"`
			TLSKey                     string `mapstructure:"tls_key"`
			JWTSecret                  string `mapstructure:"jwt_secret"`
			CORSAllowOrigin            string `mapstructure:"cors_allow_origin"`
			DisableAssignExistingUsers bool   `mapstructure:"disable_assign_existing_users"`
		} `mapstructure:"http_server"`
	} `mapstructure:"application_server"`
}

// C holds the global configuration.
var Cstruct MxpConfig
