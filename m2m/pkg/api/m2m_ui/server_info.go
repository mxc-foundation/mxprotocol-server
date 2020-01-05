package ui

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	api "github.com/mxc-foundation/mxprotocol-server/m2m/api/m2m_ui"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/config"
)

type ServerInfoAPI struct {
}

func NewServerInfoAPI() *ServerInfoAPI {
	return &ServerInfoAPI{}
}

func (s *ServerInfoAPI) GetVersion(ctx context.Context, req *empty.Empty) (*api.GetVersionResponse, error) {
	return &api.GetVersionResponse{Version: config.Cstruct.Version}, nil
}
