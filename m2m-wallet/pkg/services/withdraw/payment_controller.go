package withdraw

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/api"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/config"
	"google.golang.org/grpc"
)

func paymentReq(ctx context.Context, conf *config.MxpConfig, amount string) ( *api.TxReqReplyType, error) {
	address := conf.PaymentServer.PaymentServiceAddress + conf.PaymentServer.PaymentServicePort

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "cannot reach payment service")
	}
	defer conn.Close()

	client := api.NewPaymentClient(conn)

	//ToDo: set the correct para
	reply, err := client.TokenTxReq(ctx, &api.TxReqType{PaymentClientEnum:3,ReqIdClient:1,ReceiverAdr:"Read from DB!",
	Amount:amount,TokenNameEnum:0})
	if err != nil {
		return nil, errors.Wrap(err, "cannot get reply from payment service")
	}

	return reply, nil
}
