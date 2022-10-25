module bitbucket.org/leapllc/rep1

go 1.18

replace github.com/radiochild/repmeta => ../repmeta

require (
	github.com/aws/aws-lambda-go v1.34.1
	github.com/awslabs/aws-lambda-go-api-proxy v0.13.3
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/jackc/pgx v3.6.2+incompatible
	github.com/jackc/pgx/v4 v4.17.0
	github.com/radiochild/repmeta v0.1.0
	github.com/radiochild/utils v0.1.1
	go.uber.org/zap v1.23.0
)

require (
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.13.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.12.0 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	golang.org/x/text v0.3.7 // indirect
)
