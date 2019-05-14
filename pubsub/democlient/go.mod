module github.com/fonero-project/fnodata/pubsub/democlient

replace (
	github.com/fonero-project/fnodata/api/types => ../../api/types
	github.com/fonero-project/fnodata/blockdata => ../../blockdata
	github.com/fonero-project/fnodata/db/cache => ../../db/cache
	github.com/fonero-project/fnodata/db/dbtypes => ../../db/dbtypes
	github.com/fonero-project/fnodata/db/fnopg => ../../db/fnopg
	github.com/fonero-project/fnodata/db/fnosqlite => ../../db/fnosqlite
	github.com/fonero-project/fnodata/fnorates => ../../fnorates
	github.com/fonero-project/fnodata/exchanges => ../../exchanges
	github.com/fonero-project/fnodata/explorer/types => ../../explorer/types
	github.com/fonero-project/fnodata/gov/agendas => ../../gov/agendas
	github.com/fonero-project/fnodata/gov/politeia => ../../gov/politeia
	github.com/fonero-project/fnodata/mempool => ../../mempool
	github.com/fonero-project/fnodata/middleware => ../../middleware
	github.com/fonero-project/fnodata/pubsub => ../
	github.com/fonero-project/fnodata/pubsub/types => ../types
	github.com/fonero-project/fnodata/rpcutils => ../../rpcutils
	github.com/fonero-project/fnodata/semver => ../../semver
	github.com/fonero-project/fnodata/stakedb => ../../stakedb
	github.com/fonero-project/fnodata/testutil/dbconfig => ../../testutil/dbconfig
	github.com/fonero-project/fnodata/txhelpers => ../../txhelpers
)

require (
	github.com/google/go-cmp v0.3.0 // indirect
	github.com/jessevdk/go-flags v1.4.0
	github.com/kr/pty v1.1.4 // indirect
	golang.org/x/crypto v0.0.0-20190426145343-a29dc8fdc734 // indirect
	golang.org/x/net v0.0.0-20190502183928-7f726cade0ab
	golang.org/x/sys v0.0.0-20180926141714-2f1df4e56cde // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20190502173448-54afdca5d873 // indirect
	google.golang.org/grpc v1.20.1 // indirect
	gopkg.in/AlecAivazis/survey.v1 v1.8.2
)
