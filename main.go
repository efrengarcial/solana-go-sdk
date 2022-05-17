package main

import (
	"context"
	"fmt"
	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/metaplex/tokenvault"
	"github.com/portto/solana-go-sdk/program/sysprog"
	"github.com/portto/solana-go-sdk/program/tokenprog"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/portto/solana-go-sdk/types"
	"log"
	"time"
)

var feePayer, _ = types.AccountFromBase58("5j2o3qpvqYD8HgZM24GiXB3g3dFzmAs2rUe68se7YpUcBXcj72DsoiDres22uMZwHMPn8x17LKyNSvqaWmQAFLs1") //local-dev

func main() {

	c := client.NewClient(rpc.DevnetRPCEndpoint)
	externalPriceAccount := createExternalPriceAccount(c)

	accountRent, err := c.GetMinimumBalanceForRentExemption(context.Background(), tokenvault.AccountSize)
	if err != nil {
		log.Fatalf("failed to get mint account rent, err: %v", err)
	}

	mintRent, err := c.GetMinimumBalanceForRentExemption(context.Background(), tokenvault.MintSize)
	if err != nil {
		log.Fatalf("failed to get mint account rent, err: %v", err)
	}

	vaultRent, err := c.GetMinimumBalanceForRentExemption(context.Background(), tokenvault.MaxVaultSize)
	if err != nil {
		log.Fatalf("failed to get mint account rent, err: %v", err)
	}

	vault := types.NewAccount()
	fmt.Printf("vault: %v\n", vault.PublicKey.ToBase58())

	vaultAuthority, err := tokenvault.GetPdaForVault(vault.PublicKey)
	if err != nil {
		log.Fatalf("failed to find a valid vault authority, err: %v", err)
	}

	fractionMint := types.NewAccount()
	fmt.Printf("fractionMint: %v\n", fractionMint.PublicKey.ToBase58())

	redeemTreasury := types.NewAccount()
	fmt.Println("redeemTreasury:", redeemTreasury.PublicKey.ToBase58())

	fractionTreasury := types.NewAccount()
	fmt.Printf("fractionTreasury: %v\n", fractionTreasury.PublicKey.ToBase58())

	res, err := c.GetRecentBlockhash(context.Background())
	if err != nil {
		log.Fatalf("get recent block hash error, err: %v\n", err)
	}

	tx, err := types.NewTransaction(types.NewTransactionParam{
		Signers: []types.Account{feePayer, fractionMint, redeemTreasury, fractionTreasury, vault},
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayer.PublicKey,
			RecentBlockhash: res.Blockhash,
			Instructions: []types.Instruction{
				sysprog.CreateAccount(sysprog.CreateAccountParam{
					From:     feePayer.PublicKey,
					New:      fractionMint.PublicKey,
					Owner:    common.TokenProgramID,
					Lamports: mintRent,
					Space:    tokenprog.MintAccountSize,
				}),
				tokenprog.InitializeMint(tokenprog.InitializeMintParam{
					Decimals:   0,
					Mint:       fractionMint.PublicKey,
					MintAuth:   vaultAuthority,
					FreezeAuth: &vaultAuthority,
				}),
				sysprog.CreateAccount(sysprog.CreateAccountParam{
					From:     feePayer.PublicKey,
					New:      redeemTreasury.PublicKey,
					Owner:    common.TokenProgramID,
					Lamports: accountRent,
					Space:    tokenvault.AccountSize,
				}),
				tokenprog.InitializeAccount(tokenprog.InitializeAccountParam{
					Account: redeemTreasury.PublicKey,
					Mint:    tokenvault.QuoteMint,
					Owner:   vaultAuthority,
				}),
				sysprog.CreateAccount(sysprog.CreateAccountParam{
					From:     feePayer.PublicKey,
					New:      fractionTreasury.PublicKey,
					Owner:    common.TokenProgramID,
					Lamports: accountRent,
					Space:    tokenvault.AccountSize,
				}),
				tokenprog.InitializeAccount(tokenprog.InitializeAccountParam{
					Account: fractionTreasury.PublicKey,
					Mint:    fractionMint.PublicKey,
					Owner:   vaultAuthority,
				}),
				sysprog.CreateAccount(sysprog.CreateAccountParam{
					From:     feePayer.PublicKey,
					New:      vault.PublicKey,
					Owner:    common.MetaplexVaultProgramID,
					Lamports: vaultRent,
					Space:    tokenvault.MaxVaultSize,
				}),
				tokenvault.InitVault(tokenvault.InitVaultParam{
					Vault:                     vault.PublicKey,
					VaultAuthority:            vaultAuthority,
					FractionMint:              fractionMint.PublicKey,
					RedeemTreasury:            redeemTreasury.PublicKey,
					FractionTreasury:          fractionTreasury.PublicKey,
					PricingLookupAddress:      externalPriceAccount,
					AllowFurtherShareCreation: true,
				}),
			},
		}),
	})
	if err != nil {
		log.Fatalf("generate tx error, err: %v\n", err)
	}

	txhash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Println("txhash:", txhash)
		log.Fatalf("send tx error, err: %v\n", err)
	}

	log.Println("txhash:", txhash)
}

func createExternalPriceAccount(c *client.Client) common.PublicKey {

	epaRentExempt, err := c.GetMinimumBalanceForRentExemption(context.Background(), tokenvault.MaxExternalAccountSize)
	if err != nil {
		log.Fatalf("failed to get mint account rent, err: %v", err)
	}

	externalPriceAccount := types.NewAccount()
	fmt.Printf("ExternalPriceAccount: %v\n", externalPriceAccount.PublicKey.ToBase58())

	res, err := c.GetRecentBlockhash(context.Background())
	if err != nil {
		log.Fatalf("get recent block hash error, err: %v\n", err)
	}

	tx, err := types.NewTransaction(types.NewTransactionParam{
		Signers: []types.Account{feePayer, externalPriceAccount},
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayer.PublicKey,
			RecentBlockhash: res.Blockhash,
			Instructions: []types.Instruction{
				sysprog.CreateAccount(sysprog.CreateAccountParam{
					From:     feePayer.PublicKey,
					New:      externalPriceAccount.PublicKey,
					Owner:    common.MetaplexVaultProgramID,
					Lamports: epaRentExempt,
					Space:    tokenvault.MaxExternalAccountSize,
				}),
				tokenvault.UpdateExternalPriceAccount(tokenvault.UpdateExternalPriceAccountParam{
					PricePerShare:        0,
					PriceMint:            tokenvault.QuoteMint,
					AllowedToCombine:     true,
					ExternalPriceAccount: externalPriceAccount.PublicKey,
				}),
			},
		}),
	})
	if err != nil {
		log.Fatalf("generate tx error, err: %v\n", err)
	}

	txhash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatalf("send tx error, err: %v\n", err)
	}

	err = getTransaction(context.Background(), c, txhash)
	if err != nil {
		log.Fatalf("Get tx error, err: %v\n", err)
	}

	log.Println("txhash:", txhash)

	return externalPriceAccount.PublicKey
}

func getTransaction(ctx context.Context, c *client.Client, hash string) error {
sendtx:
	tr, err := c.GetTransaction(ctx, hash)
	if err != nil {
		return err
	}
	if tr == nil {
		time.Sleep(1 * time.Second)
		goto sendtx
	}
	return nil
}
