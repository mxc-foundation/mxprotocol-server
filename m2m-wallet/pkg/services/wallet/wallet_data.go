package wallet

type PaymentCategory string

const (
	UPLINK                PaymentCategory = "uplink"
	DOWNLINK              PaymentCategory = "downlink"
	PURCHASE_SUBSCRIPTION PaymentCategory = "purchase subscription"
	TOP_UP                PaymentCategory = "top up"
	WITHDRAW              PaymentCategory = "withdraw"
)

type DeviceType string

const (
	node    DeviceType = "node"
	gateway DeviceType = "gateway"
)

type Operation func(a float64, b float64) float64
type operMapType struct {
	pc        PaymentCategory
	dt        DeviceType
	operation Operation
}

var operMap = []operMapType{
	{UPLINK, node, uplinkNodePayment},
	{UPLINK, gateway, uplinkGatewayPayment},
	{DOWNLINK, node, downlinkNodePayment},
	{DOWNLINK, gateway, downlinkGatewayPayment},
	{PURCHASE_SUBSCRIPTION, node, psNodePayment},
	{PURCHASE_SUBSCRIPTION, gateway, psGatewayPayment},
	{TOP_UP, node, topup},
	{WITHDRAW, node, withdraw},
	{TOP_UP, gateway, topup},
	{WITHDRAW, gateway, withdraw},
}

func uplinkNodePayment(b float64, p float64) float64 {
	b = b - p
	return b
}

func uplinkGatewayPayment(b float64, p float64) float64 {
	b = b + p
	return b
}

func downlinkNodePayment(b float64, p float64) float64 {
	b = b - p
	return b
}

func downlinkGatewayPayment(b float64, p float64) float64 {
	b = b + p
	return b
}

func psNodePayment(b float64, p float64) float64 {
	b = b - p
	return b
}

func psGatewayPayment(b float64, p float64) float64 {
	b = b + p
	return b
}

func topup(b float64, p float64) float64 {
	b = b + p
	return b
}

func withdraw(b float64, p float64) float64 {
	b = b - p
	return b
}
