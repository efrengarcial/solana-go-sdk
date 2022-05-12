package fixedpricesale

import (
	"github.com/near/borsh-go"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/types"
)

type Instruction [8]uint8

var (
	InstructionCreateStore Instruction = [...]uint8{132, 152, 9, 27, 112, 19, 95, 83}
)

type CreateStoreParam struct {
	Name        string
	Description string
	Admin       common.PublicKey
	Store       common.PublicKey
}

func CreateStore(param CreateStoreParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Name        string
		Description string
	}{
		Instruction: InstructionCreateStore,
		Name:        param.Name,
		Description: param.Description,
	})
	if err != nil {
		panic(err)
	}
	return types.Instruction{
		ProgramID: common.MetaplexFixedPriceSaleProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Admin,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.Store,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     common.SystemProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}
