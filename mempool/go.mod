module github.com/fonero-project/fnodata/mempool

go 1.11

replace (
	github.com/fonero-project/fnodata/pubsub/types => ../pubsub/types
	github.com/fonero-project/fnodata/txhelpers => ../txhelpers
)

require (
	github.com/decred/slog v1.0.0
	github.com/dustin/go-humanize v1.0.0
)
