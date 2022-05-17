package tokenvault

import "github.com/portto/solana-go-sdk/common"

const (
	MaxExternalAccountSize = 1 + 8 + 32 + 1
	MaxVaultSize           = 1 + 32 + 32 + 32 + 32 + 1 + 32 + 1 + 32 + 1 + 1 + 8
	MintSize               = 82
	AccountSize            = 165
)

var (
	WrappedSolMint = common.PublicKeyFromString("So11111111111111111111111111111111111111112")
	QuoteMint      = WrappedSolMint
)
