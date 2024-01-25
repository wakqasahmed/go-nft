// Lending (10% BONUS1). An NFT owner can lend (assume no interest) the NFT to any other
// person for a period of time. After the expiration of that time period the NFT is automatically
// returned to the initial owner.

package main

import (
	"fmt"
	"time"
)

type NFT struct {
	ID            string
	Holder        string
	LendingExpiry time.Time
	IsLent        bool
	InitialOwner  string
}

type NFTRegistry map[string]NFT

func initializeNFT(holderName string) NFT {
	// Generate a unique 10-symbol alphanumeric ID
	nftID := "ABC123DEF4"
	return NFT{ID: nftID, Holder: holderName, InitialOwner: holderName}
}

func main() {
	nftRegistry := make(NFTRegistry)

	// Assign NFT to a holder
	holderName := "Waqas Ahmed"
	nft := initializeNFT(holderName)
	nftRegistry[nft.ID] = nft

	// Lend NFT
	newHolderName := "Elias Iosif"
	lendNFT(nftRegistry, nft.ID, newHolderName, time.Now().Add(24*time.Hour))

	// Query by holder
	fmt.Printf("NFTs owned by %s: %v\n", holderName, getNFTsByHolder(nftRegistry, holderName))
	fmt.Printf("NFTs owned by %s: %v\n", newHolderName, getNFTsByHolder(nftRegistry, newHolderName))
}

func lendNFT(registry NFTRegistry, nftID string, borrower string, expiry time.Time) {
	if nft, exists := registry[nftID]; exists {
		if !nft.IsLent {
			nft.IsLent = true
			nft.LendingExpiry = expiry
			nft.Holder = borrower
			registry[nftID] = nft
			fmt.Printf("NFT %s lent to %s until %s\n", nftID, borrower, expiry.Format(time.RFC3339))
		} else {
			fmt.Printf("NFT %s is already lent\n", nftID)
		}
	} else {
		fmt.Printf("NFT %s not found\n", nftID)
	}
}

// Same as Step-02
func getNFTsByHolder(registry NFTRegistry, holder string) []NFT {
	var nfts []NFT
	for _, nft := range registry {
		if nft.Holder == holder {
			nfts = append(nfts, nft)
		}
	}
	return nfts
}
