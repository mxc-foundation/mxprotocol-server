package auth

import (
	"bytes"
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/tests"
	"io/ioutil"
	"net/http"
	"testing"
)

type AUTHTestSuite struct {
	suite.Suite
}

func TESTVerifyRequestViaAuthServer(t *testing.T) {
	conf := tests.GetConfig()

	Convey("Login as admin.admin", t, func() {
		requestBody, err := json.Marshal(map[string]string{
			"password": "admin",
			"username": "admin",
		})
		So(err, ShouldBeNil)

		request, err := http.NewRequest("POST", conf.General.AuthServer+"/api/internal/login", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Accept", "application/json")
		So(err, ShouldBeNil)

		res, err := http.DefaultClient.Do(request)
		defer res.Body.Close()
		So(err, ShouldBeNil)

		body, err := ioutil.ReadAll(res.Body)
		So(err, ShouldBeNil)

		// parse response
		errInfo := errStruct{}
		err = json.Unmarshal(body, &errInfo)
		So(err, ShouldBeNil)
		So(errInfo.Error, ShouldNotEqual, "")

		var output map[string]string
		err = json.Unmarshal(body, &output)
		So(err, ShouldBeNil)

	})
}
