package tokenvault

import (
	"github.com/portto/solana-go-sdk/common"
)

func GetPdaForVault(vault common.PublicKey) (common.PublicKey, error) {
	pdaVaultAccount, _, err := common.FindProgramAddress(
		[][]byte{
			[]byte("vault"),
			common.MetaplexVaultProgramID.Bytes(),
			vault.Bytes(),
		},
		common.MetaplexVaultProgramID,
	)
	if err != nil {
		return common.PublicKey{}, err
	}
	return pdaVaultAccount, nil
}
