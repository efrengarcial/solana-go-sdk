package main

import (
	"context"
	"fmt"
	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/program/metaplex/metaplex"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/portto/solana-go-sdk/types"
	"log"
)

//var feePayer, _ = types.AccountFromBase58("h5GoYY5qrJ7XcMZDQnu1vNUjh4q1Byo7vZyiBiLUYKkfh4ikTbC4yiGzdGXrETz2M8xrsVkixYtwXjUHTMccBnQ")
//var feePayer, _ = types.AccountFromBase58("41DfuvmDMbJPNVgx8dxffpcUgnP7Fq4quiSqquFVEY6ZYNKHFUtQ83qnKUvkaAVd65ySCbah8SMD1pgAqmgxo8iM") //owner1
var feePayer, _ = types.AccountFromBase58("46R3Kii528XQ63KCP5TF23YbtWzorvdkFrp9kE2CVcyhf8o43VnvQXnEVEKs8b1DU52zihcKuvAN9eVodDZTuNfd") //local-dev

func main() {

	c := client.NewClient(rpc.DevnetRPCEndpoint)

	// create an admin account
	store, err := metaplex.GetStorePubkey(feePayer.PublicKey)
	if err != nil {
		log.Fatalf("failed to find a valid token store, err: %v", err)
	}
	//store := types.NewAccount()
	fmt.Println("store:", store)

	whitelistedCreator, err := metaplex.GetWhitelistedCreator(store, feePayer.PublicKey)
	if err != nil {
		log.Fatalf("failed to find a valid white listed Creator, err: %v", err)
	}
	fmt.Println("whitelistedCreator:", whitelistedCreator)
	/*ata, _, err := common.FindAssociatedTokenAddress(feePayer.PublicKey, store.PublicKey)
	if err != nil {
		log.Fatalf("failed to find a valid ata, err: %v", err)
	}*/

	res, err := c.GetRecentBlockhash(context.Background())
	if err != nil {
		log.Fatalf("get recent block hash error, err: %v\n", err)
	}

	/*mintAccountRent, err := c.GetMinimumBalanceForRentExemption(context.Background(), tokenprog.MintAccountSize)
	if err != nil {
		log.Fatalf("failed to get mint account rent, err: %v", err)
	}*/

	tx, err := types.NewTransaction(types.NewTransactionParam{
		Signers: []types.Account{feePayer},
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayer.PublicKey,
			RecentBlockhash: res.Blockhash,
			Instructions: []types.Instruction{
				metaplex.SetStore(metaplex.SetStoreParam{
					Admin:    feePayer.PublicKey,
					Store:    store,
					Payer:    feePayer.PublicKey,
					IsPublic: true,
					//Config:      config.PublicKey,
					//SettingsUri: "https://notgoogle.com",
				}),
				metaplex.SetWhitelistedCreator(metaplex.SetWhitelistedCreatorParam{
					Admin:                 feePayer.PublicKey,
					Store:                 store,
					Payer:                 feePayer.PublicKey,
					Creator:               feePayer.PublicKey,
					WhitelistedCreatorPDA: whitelistedCreator,
					Activated:             true,
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

	log.Println("txhash:", txhash)
}
