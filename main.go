package main

import (
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	contractAddress := common.HexToAddress("0xAdf1081Bc4B5B0596BC0830d800c7EAa87DaE03A")

	// Account information of Writer/Reader
	key := `{"address":"26f57bd43c8b9737cea37da1d5283793a0e6a4a9","crypto":{"cipher":"aes-128-ctr","ciphertext":"252d251ad0c3800da78186c62f427fd6cfb6711b519a1b723c6897438ef986e2","cipherparams":{"iv":"5574b60debba87d5b54112feb600d3ed"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"89e61af14b23d8021ae477f6e5e201d021ca166c47c9121b06fce3d722dfc2af"},"mac":"9e8dbac0aa6f0ebb4cf50c3b049dd4c212c0c9a242738ad69506c87a0749bd8b"},"id":"83386624-e582-4434-8f96-76e0ceef72d9","version":3}`
	// Passphrase to authenticate the user account on ethereum
	passphrase := "testnet";

	// geth.ipc path change it according to your env
	gethPath := "/tmp/ethereum_dev_mode/geth.ipc"

	client, err := ethclient.Dial(gethPath)

	if err != nil {
		log.Fatalf("Failed to connect Ethereum client %v:", err)
	}

	log.Printf("Connection information %v", client)

	auth, err := bind.NewTransactor(strings.NewReader(key),passphrase)

	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	log.Printf("Auth: %v", auth.From)

	contract, err := bindMain(contractAddress, client, client)

	if err != nil {
		log.Fatalf("Failed to instansitate a mapper contract: %v", err)
	}

	writerSession := &MainTransactorSession{
		Contract: &MainTransactor{
			contract: contract,
		},
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: big.NewInt(3141592),
		},
	}

	transaction, err := writerSession.MapAddress(common.HexToAddress("0xAdf1081Bc4B5B0596BC0830d800c7EAa87DaE03A"))

	if err != nil {
		log.Printf("Failed to map address: %v", err)
	}

	log.Printf("Trasaction :  %v ", transaction.Hash())

}
