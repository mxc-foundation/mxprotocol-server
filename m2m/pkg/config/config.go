package config

import (
	"time"
)

// Config defines the configuration structure.
type MxpConfig struct {
	General struct {
		LogLevel               int    `mapstructure:"log_level"`
		PasswordHashIterations int    `mapstructure:"password_hash_iterations"`
		HostServer             string `mapstructure:"host_server"`
		AuthServer             string `mapstructure:"auth_server"`
		AuthUrl                string `mapstructure:"auth_url"`
	}

	UserNotification struct {
		LowBalanceThreshold float64 `mapstructure:"low_balance_threshold"`
	} `mapstructure:"user_notification"`

	SysNotification struct {
		LowBalanceThreshold float64 `mapstructure:"low_balance_threshold"`
	} `mapstructure:"system_notification"`

	Pricing struct {
		DownLinkPkgPrice float64 `mapstructure:"downlink_package_price"`
	} `mapstructure:"pricing"`

	Accounting struct {
		IntervalMin int64 `mapstructure:"interval_min"`
	} `mapstructure:"accounting"`

	PostgreSQL struct {
		DSN         string `mapstructure:"dsn"`
		Automigrate bool   `mapstructure:"automigrate"`
	}

	Redis struct {
		URL         string        `mapstructure:"url"`
		MaxIdle     int           `mapstructure:"max_idle"`
		IdleTimeout time.Duration `mapstructure:"idle_timeout"`
	}

	M2MServer struct {
		HttpServer struct {
			Bind                       string `mapstructure:"bind"`
			TLSCert                    string `mapstructure:"tls_cert"`
			TLSKey                     string `mapstructure:"tls_key"`
			JWTSecret                  string `mapstructure:"jwt_secret"`
			CORSAllowOrigin            string `mapstructure:"cors_allow_origin"`
			DisableAssignExistingUsers bool   `mapstructure:"disable_assign_existing_users"`
		} `mapstructure:"http_server"`
	} `mapstructure:"m2m_server"`

	AppServer struct {
		Server  string `mapstructure:"appserver"`
		CACert  string `mapstructure:"ca_cert"`
		TLSCert string `mapstructure:"tls_cert"`
		TLSKey  string `mapstructure:"tls_key"`
	} `mapstructure:"appserver"`

	PaymentServer struct {
		PaymentServiceAddress string `mapstructure:"payment_service_address"`
		PaymentServicePort    string `mapstructure:"payment_service_port"`
	} `mapstructure:"paymentserver"`

	SuperNode struct {
		ContractAddress     string  `mapstructure:"contract_address"`
		SuperNodeAddress    string  `mapstructure:"supernode_address"`
		APIKey              string  `mapstructure:"api_key"`
		CheckAccountSeconds int     `mapstructure:"check_account_seconds"`
		ExtCurrAbv          string  `mapstructure:"external_currency_abv"`
		TestNet             bool    `mapstructure:"ether_test_net"`
		DlPrice             float64 `mapstructure:"down_link_price"`
	} `mapstructure:"supernode"`

	Withdraw struct {
		ResendToPS  int `mapstructure:"resend_ps_time_second"`
		RecheckStat int `mapstructure:"recheck_status_time_second"`
	} `mapstructure:"withdraw"`

	Staking struct {
		StakingPercentage float64 `mapstructure:"staking_percentage"`
		StakingMinDays    int     `mapstructure:"staking_min_days"`
	} `mapstructure:"staking"`

	Version string
}

// C holds the global configuration.
var Cstruct MxpConfig
