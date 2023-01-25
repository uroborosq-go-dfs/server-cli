module server-cli

go 1.19

require (
	github.com/google/uuid v1.3.0
	github.com/urfave/cli/v2 v2.23.7
)

require (
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/jackc/pgx v3.6.2+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/uroborosq-go-dfs/models v0.0.0-20230120113109-1d73a1f01e9b // indirect
	golang.org/x/crypto v0.5.0 // indirect
	golang.org/x/text v0.6.0 // indirect
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/uroborosq-go-dfs/server v0.0.0-20230120111813-4cca7eff115b
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
)

replace (
	github.com/uroborosq-go-dfs/server => ../server
	github.com/uroborosq-go-dfs/server/connector => ../server/connector
	github.com/uroborosq-go-dfs/server/node => ../server/node
)
