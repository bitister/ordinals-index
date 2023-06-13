module syncer

go 1.20

replace utils => ../utils

replace enum => ../enum

replace models => ../models

require (
	enum v0.0.0-00010101000000-000000000000
	github.com/PuerkitoBio/goquery v1.8.1
	github.com/astaxie/beego v1.12.3
	github.com/json-iterator/go v1.1.12
	models v0.0.0-00010101000000-000000000000
	utils v0.0.0-00010101000000-000000000000
)

require (
	github.com/andybalholm/cascadia v1.3.2 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.2.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/ethereum/go-ethereum v1.11.5 // indirect
	github.com/go-redis/redis v6.15.9+incompatible // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/holiman/uint256 v1.2.0 // indirect
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mervick/aes-everywhere/go/aes256 v0.0.0-20220903070135-f13ed3789ae1 // indirect
	github.com/miguelmota/go-solidity-sha3 v0.1.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.15.1 // indirect
	github.com/prometheus/client_model v0.4.0 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/prometheus/procfs v0.10.1 // indirect
	github.com/shiena/ansicolor v0.0.0-20230509054315-a9deabde6e02 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	golang.org/x/crypto v0.9.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
