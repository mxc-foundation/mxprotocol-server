package ui

import (
	"context"
	api "github.com/mxc-foundation/mxprotocol-server/m2m/api/m2m_ui"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/config"
	"testing"
)

func TestGetDlPrice(t *testing.T) {
	s := WalletServerAPI{}
	req := &api.GetDownLinkPriceRequest{OrgId: 1}
	resp, err := s.GetDlPrice(context.Background(), req)
	if err != nil {
		t.Errorf("GetDlPriceTest got unexpected error")
	}

	if resp.DownLinkPrice != config.Cstruct.SuperNode.DlPrice {
		t.Errorf("GetDlPrice=%v, wanted %v", resp.DownLinkPrice, config.Cstruct.SuperNode.DlPrice)
	}
}
