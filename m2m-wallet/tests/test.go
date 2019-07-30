package tests

import (
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/config"
	"os"
)

func GetConfig() config.MxpConfig {
	log.SetLevel(log.ErrorLevel)
	var c config.MxpConfig

	if v := os.Getenv("TEST_POSTGRES_DSN"); v != "" {
		c.PostgreSQL.DSN = v
	}

	if v := os.Getenv("TEST_LORA_APP_SERVER"); v != "" {
		c.General.AuthServer = v
	}


	return c
}