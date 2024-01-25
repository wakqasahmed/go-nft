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
NFT ABC123DEF4 lent to Elias Iosif until 2024-01-25T15:56:53+05:00
NFTs owned by Waqas Ahmed: [{ABC123DEF4 Waqas Ahmed Elias Iosif 2024-01-25 15:56:53.102992 +0500 PKT m=+60.000136459 true}]
NFTs owned by Elias Iosif: []
Simulating 1 minute 1 second time passing
NFT ABC123DEF4 automatically returned to Waqas Ahmed
Now the time is: 2024-01-25T15:56:54+05:00
NFTs owned by Waqas Ahmed: [{ABC123DEF4 Waqas Ahmed Waqas Ahmed 0001-01-01 00:00:00 +0000 UTC false}]
NFTs owned by Elias Iosif: []
```
