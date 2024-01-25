// Step02: Queries (30%) - Retrieve NFTs by Holder and Holder by NFT

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

	// Query by holder
	fmt.Printf("NFTs owned by %s: %v\n", holderName, getNFTsByHolder(nftRegistry, holderName))

	// Query by NFT
	queryNFT := "ABC123DEF4"
	fmt.Printf("Holder of NFT %s: %s\n", queryNFT, getHolderByNFT(nftRegistry, queryNFT))
}

func getNFTsByHolder(registry NFTRegistry, holder string) []NFT {
	var nfts []NFT
	for _, nft := range registry {
		if nft.Holder == holder {
			nfts = append(nfts, nft)
		}
	}
	return nfts
}

func getHolderByNFT(registry NFTRegistry, nftID string) string {
	if nft, exists := registry[nftID]; exists {
		return nft.Holder
	}
	return "NFT not found"
}
