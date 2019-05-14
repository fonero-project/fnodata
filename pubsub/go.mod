module github.com/fonero-project/fnodata/pubsub

go 1.11

replace github.com/fonero-project/fnodata/pubsub/types => ./types

replace github.com/fonero-project/fnodata/mempool => ../mempool

require (
	github.com/DataDog/zstd v1.3.8 // indirect
	github.com/decred/slog v1.0.0
	golang.org/x/net v0.0.0-20190415214537-1da14a5a36f2
)
