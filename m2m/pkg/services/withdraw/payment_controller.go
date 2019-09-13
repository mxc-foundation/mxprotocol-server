package withdraw

import (
	"context"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	ps "gitlab.com/MXCFoundation/cloud/mxprotocol-server/grpc_api-paymemt_service"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
	"google.golang.org/grpc"
)

var PaymentServiceAvailable bool

const (
	TOBE_SENT_FROM_PAYMENT_SERVER    = "TOBE_SENT"
	TOBE_CHECKED_FROM_PAYMENT_SERVER = "TOBE_CHECKED" // tx is sent, still not sure if it was successful
	SUCCESSFUL                       = "SUCCESSFUL"
)

func paymentServiceAvailable(conf config.MxpConfig) bool {
	log.Info("/withdraw: try to connect to payment service: ",
		conf.PaymentServer.PaymentServiceAddress+conf.PaymentServer.PaymentServicePort)
	conn, err := grpc.Dial(conf.PaymentServer.PaymentServiceAddress+conf.PaymentServer.PaymentServicePort,
		grpc.WithInsecure())
	if err != nil {
		log.WithError(err).Error("/withdraw: payment service is not available.")
		return false
	}

	defer conn.Close()

	return true
}

func PaymentReq(ctx context.Context, conf *config.MxpConfig, amount, receiverAdd string, reqId int64) (*ps.TxReqReplyType, error) {
	address := conf.PaymentServer.PaymentServiceAddress + conf.PaymentServer.PaymentServicePort

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "cannot reach payment service")
	}
	defer conn.Close()

	client := ps.NewPaymentClient(conn)

	reply, err := client.TokenTxReq(ctx, &ps.TxReqType{PaymentClientEnum: 3, ReqIdClient: reqId, ReceiverAdr: receiverAdd,
		Amount: amount, TokenNameEnum: 0})
	if err != nil {
		return nil, errors.Wrap(err, "cannot get reply from payment service")
	}

	return reply, nil
}

func CheckTxStatus(conf *config.MxpConfig, qreID int64) (*ps.CheckTxStatusReplyType, error) {
	address := conf.PaymentServer.PaymentServiceAddress + conf.PaymentServer.PaymentServicePort

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "cannot reach payment service")
	}
	defer conn.Close()

	client := ps.NewPaymentClient(conn)

	//ToDo: get the ReqID from db
	reply, err := client.CheckTxStatus(context.Background(), &ps.CheckTxStatusType{ReqQueryRef: qreID})
	if err != nil {
		return nil, errors.Wrap(err, "cannot get reply from payment service")
	}

	return reply, nil
}

