package withdraw

import (
	"context"
	"github.com/pkg/errors"
	ps "gitlab.com/MXCFoundation/cloud/mxprotocol-server/grpc_api-paymemt_service"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/config"
	"google.golang.org/grpc"
)

func paymentReq(ctx context.Context, conf *config.MxpConfig, amount string) (*ps.TxReqReplyType, error) {
	address := conf.PaymentServer.PaymentServiceAddress + conf.PaymentServer.PaymentServicePort

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "cannot reach payment service")
	}
	defer conn.Close()

	client := ps.NewPaymentClient(conn)

	//ToDo: set the correct para
	reply, err := client.TokenTxReq(ctx, &ps.TxReqType{PaymentClientEnum: 3, ReqIdClient: 1, ReceiverAdr: "Read from DB!",
		Amount: amount, TokenNameEnum: 0})
	if err != nil {
		return nil, errors.Wrap(err, "cannot get reply from payment service")
	}

	return reply, nil
}
