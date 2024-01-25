# go-nft
NFT assignment Go based for Permissioned blockchain programming BLOC-523

# How to run

### Check if go is installed
```
$ go version
go version go1.20.2 darwin/arm64
```

### Build each step
```
$ go build -o ./bin ./step01-initialization.go
$ go build -o ./bin ./step02-queries.go 
$ go build -o ./bin ./step03-transfer-of-ownership.go 
$ go build -o ./bin ./step04-lending.go 
```

### Execute each step
```
$ ./bin/step01-initialization
NFT ABC123DEF4 assigned to Waqas Ahmed

$ ./bin/step02-queries 
NFTs owned by Waqas Ahmed: [{ABC123DEF4 Waqas Ahmed}]
Holder of NFT ABC123DEF4: Waqas Ahmed

$ ./bin/step03-transfer-of-ownership 
NFTs owned by Waqas Ahmed: [{ABC123DEF4 Waqas Ahmed}]
NFTs owned by Elias Iosif: []
NFT ABC123DEF4 transferred from Waqas Ahmed to Elias Iosif
NFTs owned by Waqas Ahmed: []
NFTs owned by Elias Iosif: [{ABC123DEF4 Elias Iosif}]

$ ./bin/step04-lending 
NFT ABC123DEF4 lent to Elias Iosif until 2024-01-26T15:21:31+05:00
NFTs owned by Waqas Ahmed: []
NFTs owned by Elias Iosif: [{ABC123DEF4 Elias Iosif 2024-01-26 15:21:31.884346 +0500 PKT m=+86400.000139251 true Waqas Ahmed}]
```
