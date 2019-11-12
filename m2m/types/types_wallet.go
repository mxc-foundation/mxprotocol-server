package types

type WalletType string

const (

	// used for managing topup, withdraw, aggregation of internal transactions
	SUPER_ADMIN WalletType = "SUPER_ADMIN"

	// used for income of the super node
	SUPER_NODE_INCOME WalletType = "SUPER_NODE_INCOME"

	// All the stakes will come to this account
	STAKE_STORAGE WalletType = "STAKE_STORAGE"

	// Normal users of MXProtocol
	USER WalletType = "USER"
)
