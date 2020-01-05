package services

import (
	"github.com/stretchr/testify/suite"
	"github.com/mxc-foundation/mxprotocol-server/m2m/db"
)

type DbInterfaceTestSuite struct {
	tx *db.TxHandler
}

type ServicesTestSuite struct {
	suite.Suite
	DbInterfaceTestSuite
}
