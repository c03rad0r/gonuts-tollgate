package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/Origami74/gonuts-tollgate/wallet"
)

func main() {
	// Create a temporary directory for the wallet
	tmpDir, err := os.MkdirTemp("", "wallet-bug-test")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	walletPath := filepath.Join(tmpDir, "wallet")
	config := wallet.Config{
		WalletPath:     walletPath,
		CurrentMintURL: "https://8333.space:3338",
	}

	fmt.Println("Loading wallet first time...")
	start := time.Now()
	
	// First load should work
	wallet1, err := wallet.LoadWallet(config)
	if err != nil {
		fmt.Printf("First load failed: %v\n", err)
	} else {
		fmt.Printf("First load succeeded in %v\n", time.Since(start))
		fmt.Printf("Wallet balance: %d\n", wallet1.GetBalance())
		_ = wallet1.Shutdown()
	}

	fmt.Println("Loading wallet second time...")
	start = time.Now()

	// Second load should also work but might hang
	wallet2, err := wallet.LoadWallet(config)
	if err != nil {
		fmt.Printf("Second load failed: %v\n", err)
	} else {
		fmt.Printf("Second load succeeded in %v\n", time.Since(start))
		fmt.Printf("Wallet balance: %d\n", wallet2.GetBalance())
		_ = wallet2.Shutdown()
	}
	
	fmt.Println("Test completed")
}