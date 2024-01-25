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
	Owner         string
	CurrentHolder string
	LendingExpiry time.Time
	IsLent        bool
}

type NFTRegistry map[string]NFT

func initializeNFT(ownerName string) NFT {
	// Generate a unique 10-symbol alphanumeric ID
	nftID := "ABC123DEF4"
	return NFT{ID: nftID, Owner: ownerName, CurrentHolder: ownerName}
}

func main() {
	nftRegistry := make(NFTRegistry)

	// Assign NFT to an owner
	ownerName := "Waqas Ahmed"
	nft := initializeNFT(ownerName)
	nftRegistry[nft.ID] = nft

	// Lend NFT
	borrowerName := "Elias Iosif"
	lendNFT(nftRegistry, nft.ID, borrowerName, time.Now().Add(time.Minute))

	// Check ownership after lending
	fmt.Printf("NFTs owned by %s: %v\n", ownerName, getNFTsByOwner(nftRegistry, ownerName))
	fmt.Printf("NFTs owned by %s: %v\n", borrowerName, getNFTsByOwner(nftRegistry, borrowerName))

	// Simulate time passing
	fmt.Printf("Simulating 1 minute 1 second time passing\n")
	pastTime := 61 * time.Second
	time.Sleep(pastTime)
	fmt.Printf("Now the time is: %s\n", time.Now().Format(time.RFC3339))

	// Check ownership after lending period expiration
	fmt.Printf("NFTs owned by %s: %v\n", ownerName, getNFTsByOwner(nftRegistry, ownerName))
	fmt.Printf("NFTs owned by %s: %v\n", borrowerName, getNFTsByOwner(nftRegistry, borrowerName))
}

func lendNFT(registry NFTRegistry, nftID string, borrower string, expiry time.Time) {
	if nft, exists := registry[nftID]; exists {
		if !nft.IsLent {
			owner := nft.CurrentHolder
			nft.IsLent = true
			nft.LendingExpiry = expiry
			nft.CurrentHolder = borrower
			registry[nftID] = nft
			fmt.Printf("NFT %s lent to %s until %s\n", nftID, borrower, expiry.Format(time.RFC3339))
			go returnNFTAfterExpiry(registry, nftID, owner, expiry)
		} else {
			fmt.Printf("NFT %s is already lent\n", nftID)
		}
	} else {
		fmt.Printf("NFT %s not found\n", nftID)
	}
}

func returnNFTAfterExpiry(registry NFTRegistry, nftID string, owner string, expiry time.Time) {
	// Simulate time passing
	time.Sleep(time.Until(expiry))

	if nft, exists := registry[nftID]; exists {
		if nft.IsLent && time.Now().After(expiry) {
			nft.IsLent = false
			nft.LendingExpiry = time.Time{}
			nft.CurrentHolder = owner
			registry[nftID] = nft
			fmt.Printf("NFT %s automatically returned to %s\n", nftID, owner)
		}
	}
}

func getNFTsByOwner(registry NFTRegistry, owner string) []NFT {
	var nfts []NFT
	for _, nft := range registry {
		if nft.Owner == owner {
			nfts = append(nfts, nft)
		}
	}
	return nfts
}
