package accounting

import (
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/mxc-foundation/mxprotocol-server/m2m/db"
	"github.com/mxc-foundation/mxprotocol-server/m2m/tests"
)

func TestAccounting(t *testing.T) {

	conf := tests.GetConfig()
	// conf.PostgreSQL.DSN should be defined for test in docker-compose.yml

	err := db.Setup(conf)
	assertions.ShouldBeNil(err)

	if err := performAccounting(60, 0.01); err != nil {
		t.Log("error: ", err)
	}

	assertions.ShouldBeNil(err)

}
