package services

import (
	"github.com/stretchr/testify/suite"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db"
)

type DbInterfaceTestSuite struct {
	tx *db.TxHandler
}

type ServicesTestSuite struct {
	suite.Suite
	DbInterfaceTestSuite
}
