// Transfer of ownership (40%). The ownership of an NFT can be transferred between two
// entities (from owner to any other person).

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
	holder1Name := "Waqas Ahmed"
	nft := initializeNFT(holder1Name)
	nftRegistry[nft.ID] = nft

	// New NFT holder
	holder2Name := "Elias Iosif"

	// Query by holder before transfer of ownership
	fmt.Printf("NFTs owned by %s: %v\n", holder1Name, getNFTsByHolder(nftRegistry, holder1Name))
	fmt.Printf("NFTs owned by %s: %v\n", holder2Name, getNFTsByHolder(nftRegistry, holder2Name))

	// Transfer ownership
	transferNFT(nftRegistry, nft.ID, holder2Name)

	// Query by holder after transfer of ownership
	fmt.Printf("NFTs owned by %s: %v\n", holder1Name, getNFTsByHolder(nftRegistry, holder1Name))
	fmt.Printf("NFTs owned by %s: %v\n", holder2Name, getNFTsByHolder(nftRegistry, holder2Name))
}

func transferNFT(registry NFTRegistry, nftID string, newHolder string) {
	if nft, exists := registry[nftID]; exists {
		var oldHolder = nft.Holder
		nft.Holder = newHolder
		registry[nftID] = nft
		fmt.Printf("NFT %s transferred from %s to %s\n", nftID, oldHolder, nft.Holder)
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
