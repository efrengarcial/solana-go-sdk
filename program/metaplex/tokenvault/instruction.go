package tokenvault

import (
	"github.com/near/borsh-go"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/types"
)

type Key borsh.Enum

const (
	Uninitialized Key = iota
	VaultV1
	SafetyDepositBoxV1
	ExternalPriceAccountV1
)

type Instruction uint8

const (
	InstructionInitVault                  Instruction = 0
	InstructionUpdateExternalPriceAccount Instruction = 9
)

type UpdateExternalPriceAccountParam struct {
	Key                  Key
	PricePerShare        uint64
	PriceMint            common.PublicKey
	AllowedToCombine     bool
	ExternalPriceAccount common.PublicKey
}

func UpdateExternalPriceAccount(param UpdateExternalPriceAccountParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction      Instruction
		Key              Key
		PricePerShare    uint64
		PriceMint        common.PublicKey
		AllowedToCombine bool
	}{
		Instruction:      InstructionUpdateExternalPriceAccount,
		Key:              ExternalPriceAccountV1,
		PricePerShare:    param.PricePerShare,
		PriceMint:        param.PriceMint,
		AllowedToCombine: param.AllowedToCombine,
	})
	if err != nil {
		panic(err)
	}
	return types.Instruction{
		ProgramID: common.MetaplexVaultProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.ExternalPriceAccount,
				IsSigner:   false,
				IsWritable: true,
			},
		},
		Data: data,
	}
}

type InitVaultParam struct {
	Vault                     common.PublicKey
	VaultAuthority            common.PublicKey
	FractionMint              common.PublicKey
	RedeemTreasury            common.PublicKey
	FractionTreasury          common.PublicKey
	PricingLookupAddress      common.PublicKey
	AllowFurtherShareCreation bool
}

func InitVault(param InitVaultParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction               Instruction
		AllowFurtherShareCreation bool
	}{
		Instruction:               InstructionInitVault,
		AllowFurtherShareCreation: param.AllowFurtherShareCreation,
	})
	if err != nil {
		panic(err)
	}
	return types.Instruction{
		ProgramID: common.MetaplexVaultProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.FractionMint,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.RedeemTreasury,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.FractionTreasury,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Vault,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.VaultAuthority,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.PricingLookupAddress,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.TokenProgramID,
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
