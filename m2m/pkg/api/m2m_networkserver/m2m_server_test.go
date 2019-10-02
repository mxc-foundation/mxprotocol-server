package m2m_networkserver

import (
	"context"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/networkserver"
	"testing"
)

func TestDvUsageMode(t *testing.T) {
	//modes := []string{"DV_INACTIVE", "DV_FREE_GATEWAYS_LIMITED", "DV_WHOLE_NETWORK", "DV_uDELETED"}
	s := M2MNetworkServerAPI{}

	//expected := networkserver.DvUsageModeResponse{}
	req := &networkserver.DvUsageModeRequest{DvEui: "dv1"}
	resp, err := s.DvUsageMode(context.Background(), req)
	if err != nil {
		t.Errorf("Test got unexpected error %s", err)
	}

	if resp.DvMode.String() != "DV_FREE_GATEWAYS_LIMITED" {
		//t.Errorf("Expected req %s to be %s but instead got %s!", req, expected, actual)
		t.Errorf("Expected not")
	}

}
