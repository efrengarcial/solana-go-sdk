package metaplex

import (
	"github.com/near/borsh-go"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/types"
)

type Instruction uint8

const (
	InstructionSetStore              Instruction = 8
	InstructionSetWhitelistedCreator Instruction = 9
	InstructionSetStoreV2            Instruction = 23
)

type SetStoreV2Param struct {
	Admin       common.PublicKey
	Store       common.PublicKey
	Config      common.PublicKey
	Payer       common.PublicKey
	IsPublic    bool
	SettingsUri string
}

func SetStoreV2(param SetStoreV2Param) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Public      bool
		SettingsUri *string
	}{
		Instruction: InstructionSetStoreV2,
		Public:      param.IsPublic,
		SettingsUri: &param.SettingsUri,
	})
	if err != nil {
		panic(err)
	}
	return types.Instruction{
		ProgramID: common.MetaplexMetaplexProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Store,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Config,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Admin,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.Payer,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     common.TokenProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.MetaplexVaultProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.MetaplexTokenMetaProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.MetaplexAuctionProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SystemProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SysVarRentPubkey,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

type SetStoreParam struct {
	Admin    common.PublicKey
	Store    common.PublicKey
	Payer    common.PublicKey
	IsPublic bool
}

func SetStore(param SetStoreParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Public      bool
	}{
		Instruction: InstructionSetStore,
		Public:      param.IsPublic,
	})
	if err != nil {
		panic(err)
	}
	return types.Instruction{
		ProgramID: common.MetaplexMetaplexProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Store,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Admin,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.Payer,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     common.TokenProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.MetaplexVaultProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.MetaplexTokenMetaProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.MetaplexAuctionProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SystemProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SysVarRentPubkey,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

type SetWhitelistedCreatorParam struct {
	Admin                 common.PublicKey
	Store                 common.PublicKey
	Payer                 common.PublicKey
	WhitelistedCreatorPDA common.PublicKey
	Creator               common.PublicKey
	Activated             bool
}

func SetWhitelistedCreator(param SetWhitelistedCreatorParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Activated   bool
	}{
		Instruction: InstructionSetWhitelistedCreator,
		Activated:   param.Activated,
	})
	if err != nil {
		panic(err)
	}
	return types.Instruction{
		ProgramID: common.MetaplexMetaplexProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.WhitelistedCreatorPDA,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Admin,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.Payer,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.Creator,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.Store,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SystemProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SysVarRentPubkey,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}
