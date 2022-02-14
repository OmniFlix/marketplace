module github.com/OmniFlix/marketplace

go 1.16

require (
	github.com/OmniFlix/onft v0.3.0-alpha.1
	github.com/cosmos/cosmos-sdk v0.45.1
	github.com/cosmos/ibc-go v1.2.2
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.2.0
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	github.com/tendermint/spm v0.1.9
	github.com/tendermint/tendermint v0.34.15
	github.com/tendermint/tm-db v0.6.6
	google.golang.org/genproto v0.0.0-20210917145530-b395a37504d4
	google.golang.org/grpc v1.42.0
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
