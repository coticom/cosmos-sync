package msgparser

const (
	BankRouteKey         string = "bank"
	StakingRouteKey      string = "staking"
	DistributionRouteKey string = "distribution"
	CrisisRouteKey       string = "crisis"
	EvidenceRouteKey     string = "evidence"
	GovRouteKey          string = "gov"
	SlashingRouteKey     string = "slashing"
	NftRouteKey          string = "nft"
	ServiceRouteKey      string = "service"
	TokenRouteKey        string = "token"
	HtlcRouteKey         string = "htlc"
	CoinswapRouteKey     string = "coinswap"
	RandomRouteKey       string = "random"
	OracleRouteKey       string = "oracle"
	RecordRouteKey       string = "record"
)

var RouteHandlerMap = map[string]Handler{
	BankRouteKey:         handleBank,
	StakingRouteKey:      handleStaking,
	DistributionRouteKey: handleDistribution,
	CrisisRouteKey:       handleCrisis,
	EvidenceRouteKey:     handleEvidence,
	GovRouteKey:          handleGov,
	SlashingRouteKey:     handleSlashing,
	NftRouteKey:          handleNft,
	ServiceRouteKey:      handleService,
	TokenRouteKey:        handleToken,
	HtlcRouteKey:         handleHtlc,
	CoinswapRouteKey:     handleCoinswap,
	RandomRouteKey:       handleRandom,
	OracleRouteKey:       handleOracle,
	RecordRouteKey:       handleRecord,
}