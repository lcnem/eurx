package types

const (
	// ModuleName defines the module name
	ModuleName = "incentive"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	ParamsKey = "Params-value-"

	EurxMintingRewardDenom = "uestm"
)

var (
	EurxMintingClaimKeyPrefix                     = []byte{0x01} // prefix for keys that store Eurx minting claims
	EurxMintingRewardFactorKeyPrefix              = []byte{0x02} // prefix for key that stores Eurx minting reward factors
	PreviousEurxMintingRewardAccrualTimeKeyPrefix = []byte{0x03} // prefix for key that stores the blocktime
)
