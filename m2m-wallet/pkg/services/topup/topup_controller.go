package topup

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/api"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Setup() error {
	log.Info("Setup top_up service")
	return nil
}

type TopUpServerAPI struct {
	serviceName string
}

func NewTopUpServerAPI() *TopUpServerAPI {
	return &TopUpServerAPI{serviceName: "top up"}
}

func (s *TopUpServerAPI) GetTopUpHistory(ctx context.Context, req *api.GetTopUpHistoryRequest) (*api.GetTopUpHistoryResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.JsonParseError:
	case auth.ErrorInfoNotNull:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdDeleted:
		return &api.GetTopUpHistoryResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")

	case auth.OK:

		log.WithFields(log.Fields{
			"orgId": req.OrgId,
			"offset": req.Offset,
			"limit": req.Limit,
		}).Debug("grpc_api/GetTopUpHistory")

		return &api.GetTopUpHistoryResponse{UserProfile: &userProfile}, nil

	}

	return nil, status.Errorf(codes.Unknown, "")
}
