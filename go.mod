module github.com/fonero-project/fnodata

require (
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/chappjc/logrus-prefix v0.0.0-20180227015900-3a1d64819adb
	github.com/decred/slog v1.0.0
	github.com/dgryski/go-farm v0.0.0-20190416075124-e1214b5e05dc // indirect
	github.com/didip/tollbooth v4.0.1-0.20180415195142-b10a036da5f0+incompatible
	github.com/didip/tollbooth_chi v0.0.0-20170928041846-6ab5f3083f3d
	github.com/dmigwi/go-piparser/proposals v0.0.0-20190426030541-8412e0f44f55
	github.com/dustin/go-humanize v1.0.0
	github.com/go-chi/chi v4.0.3-0.20190316151245-d08916613452+incompatible
	github.com/google/gops v0.3.5
	github.com/googollee/go-engine.io v0.0.0-20180829091931-e2f255711dcb // indirect
	github.com/googollee/go-socket.io v0.0.0-20181214084611-0ad7206c347a
	github.com/jessevdk/go-flags v1.4.0
	github.com/jrick/logrotate v1.0.0
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/kisielk/gotool v1.0.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/rs/cors v1.6.0
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644
	github.com/sirupsen/logrus v1.2.0
	github.com/smartystreets/goconvey v0.0.0-20190330032615-68dc04aab96a // indirect
	github.com/x-cray/logrus-prefixed-formatter v0.5.2 // indirect
	golang.org/x/net v0.0.0-20190415214537-1da14a5a36f2
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 // indirect
)

replace (
	github.com/fonero-project/fnodata/api/types => ./api/types
	github.com/fonero-project/fnodata/blockdata => ./blockdata
	github.com/fonero-project/fnodata/db/cache => ./db/cache
	github.com/fonero-project/fnodata/db/dbtypes => ./db/dbtypes
	github.com/fonero-project/fnodata/db/fnopg => ./db/fnopg
	github.com/fonero-project/fnodata/db/fnosqlite => ./db/fnosqlite
	github.com/fonero-project/fnodata/exchanges => ./exchanges
	github.com/fonero-project/fnodata/explorer/types => ./explorer/types
	github.com/fonero-project/fnodata/fnorates => ./fnorates
	github.com/fonero-project/fnodata/gov/agendas => ./gov/agendas
	github.com/fonero-project/fnodata/gov/politeia => ./gov/politeia
	github.com/fonero-project/fnodata/mempool => ./mempool
	github.com/fonero-project/fnodata/middleware => ./middleware
	github.com/fonero-project/fnodata/pubsub => ./pubsub
	github.com/fonero-project/fnodata/pubsub/types => ./pubsub/types
	github.com/fonero-project/fnodata/rpcutils => ./rpcutils
	github.com/fonero-project/fnodata/semver => ./semver
	github.com/fonero-project/fnodata/stakedb => ./stakedb
	github.com/fonero-project/fnodata/testutil/dbconfig => ./testutil/dbconfig
	github.com/fonero-project/fnodata/txhelpers => ./txhelpers
)
