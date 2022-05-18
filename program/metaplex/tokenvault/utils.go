package tokenvault

import (
	"github.com/portto/solana-go-sdk/common"
)

func GetPdaForVault(vault common.PublicKey) (common.PublicKey, error) {
	pda, _, err := common.FindProgramAddress(
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
	return pda, nil
}

func GetSafetyDepositAccount(vault, tokenMint common.PublicKey) (common.PublicKey, error) {
	pda, _, err := common.FindProgramAddress(
		[][]byte{
			[]byte("vault"),
			vault.Bytes(),
			tokenMint.Bytes(),
		},
		common.MetaplexVaultProgramID,
	)
	if err != nil {
		return common.PublicKey{}, err
	}
	return pda, nil
}
