package msgparser

const (
	BankRouteKey         string = "bank"
	StakingRouteKey      string = "staking"
	DistributionRouteKey string = "distribution"
	CrisisRouteKey       string = "crisis"
	EvidenceRouteKey     string = "evidence"
	GovRouteKey          string = "gov"
	FeegrantRouteKey     string = "feegrant"
	SlashingRouteKey     string = "slashing"
	NftRouteKey          string = "nft"
	MtRouteKey           string = "mt"
	ServiceRouteKey      string = "service"
	TokenRouteKey        string = "token"
	HtlcRouteKey         string = "htlc"
	CoinswapRouteKey     string = "coinswap"
	RandomRouteKey       string = "random"
	OracleRouteKey       string = "oracle"
	IdentityRouteKey     string = "identity"
	RecordRouteKey       string = "record"
	IbcRouteKey          string = "ibc"
	IbcTransferRouteKey  string = "transfer"
)

var RouteHandlerMap = map[string]Handler{
	BankRouteKey:         handleBank,
	StakingRouteKey:      handleStaking,
	DistributionRouteKey: handleDistribution,
	CrisisRouteKey:       handleCrisis,
	EvidenceRouteKey:     handleEvidence,
	GovRouteKey:          handleGov,
	FeegrantRouteKey:     handleFeegrant,
	SlashingRouteKey:     handleSlashing,
	NftRouteKey:          handleNft,
	MtRouteKey:           handleMt,
	ServiceRouteKey:      handleService,
	TokenRouteKey:        handleToken,
	HtlcRouteKey:         handleHtlc,
	CoinswapRouteKey:     handleCoinswap,
	RandomRouteKey:       handleRandom,
	OracleRouteKey:       handleOracle,
	IdentityRouteKey:     handleIdentity,
	RecordRouteKey:       handleRecord,
	IbcRouteKey:          handleIbc,
	IbcTransferRouteKey:  handleIbc,
}
