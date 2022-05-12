package metaplex

import (
	"github.com/portto/solana-go-sdk/common"
)

func GetStorePubkey(payer common.PublicKey) (common.PublicKey, error) {
	storeAccount, _, err := common.FindProgramAddress(
		[][]byte{
			[]byte("metaplex"),
			common.MetaplexMetaplexProgramID.Bytes(),
			payer.Bytes(),
		},
		common.MetaplexMetaplexProgramID,
	)
	if err != nil {
		return common.PublicKey{}, err
	}
	return storeAccount, nil
}

func GetWhitelistedCreator(store common.PublicKey, creator common.PublicKey) (common.PublicKey, error) {
	whitelistedCreator, _, err := common.FindProgramAddress(
		[][]byte{
			[]byte("metaplex"),
			common.MetaplexMetaplexProgramID.Bytes(),
			store.Bytes(),
			creator.Bytes(),
		},
		common.MetaplexMetaplexProgramID,
	)
	if err != nil {
		return common.PublicKey{}, err
	}
	return whitelistedCreator, nil
}
