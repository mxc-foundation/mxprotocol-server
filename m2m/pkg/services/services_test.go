package services

import (
	"github.com/stretchr/testify/suite"
	"github.com/mxc-foundation/mxprotocol-server/m2m/db"
	"github.com/mxc-foundation/mxprotocol-server/m2m/tests"
	"testing"
)

// SetupSuite is called once before starting the test-suite.
func (b *DbInterfaceTestSuite) SetupSuite() {
	conf := tests.GetConfig()
	if err := db.Setup(conf); err != nil {
		panic(err)
	}
}

// SetupTest is called before every test.
func (b *DbInterfaceTestSuite) SetupTest() {
	tx, err := db.DB().Begin()
	if err != nil {
		panic(err)
	}
	b.tx = tx

	tests.ResetDB(db.DB().DB)
}

// TearDownTest is called after every test.
func (b *DbInterfaceTestSuite) TearDownTest() {
	if err := b.tx.Rollback(); err != nil {
		panic(err)
	}
}

// Tx returns a database transaction (which is rolled back after every
// test).
func (b *DbInterfaceTestSuite) Tx() *db.TxHandler {
	return b.tx
}

func TestServices(t *testing.T) {
	suite.Run(t, new(ServicesTestSuite))
}
