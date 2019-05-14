module github.com/fonero-project/fnodata/db/fnosqlite

go 1.12

replace github.com/fonero-project/fnodata/stakedb => ../../stakedb

require (
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/decred/slog v1.0.0
	github.com/dustin/go-humanize v1.0.0
	github.com/google/go-cmp v0.2.0
	github.com/mattn/go-sqlite3 v1.10.0
)

replace (
	github.com/fonero-project/fnodata/db/cache => ../cache
	github.com/fonero-project/fnodata/testutil/dbconfig => ../../testutil/dbconfig
)
