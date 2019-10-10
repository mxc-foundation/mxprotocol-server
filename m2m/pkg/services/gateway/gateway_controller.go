package gateway

import (
	log "github.com/sirupsen/logrus"
)

func Setup() error {
	log.Info("Setup gateway service")
	return nil
}

func syncGatewaysFromAppserver()(error) {
	// call api from appserver to update gateway list

	return nil
}