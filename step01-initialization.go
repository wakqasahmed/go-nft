// Step01: Initialization (30%). An NFT can be assigned to a holder.

package main

import (
	"fmt"
)

type NFT struct {
	ID     string
	Holder string
}

type NFTRegistry map[string]NFT

func initializeNFT(holderName string) NFT {
	// Generate a unique 10-symbol alphanumeric ID
	nftID := "ABC123DEF4"
	return NFT{ID: nftID, Holder: holderName}
}

func main() {
	nftRegistry := make(NFTRegistry)

	// Assign NFT to a holder
	holderName := "Waqas Ahmed"
	nft := initializeNFT(holderName)
	nftRegistry[nft.ID] = nft

	fmt.Printf("NFT %s assigned to %s\n", nft.ID, nft.Holder)
}
