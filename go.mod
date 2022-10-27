module github.com/radiochild/rep1

go 1.18

replace github.com/radiochild/repmeta => ../repmeta

require (
	github.com/aws/aws-lambda-go v1.34.1
	github.com/aws/aws-sdk-go-v2 v1.17.1
	github.com/aws/aws-sdk-go-v2/config v1.17.10
	github.com/aws/aws-sdk-go-v2/service/s3 v1.29.1
	github.com/aws/smithy-go v1.13.4
	github.com/awslabs/aws-lambda-go-api-proxy v0.13.3
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/jackc/pgx/v4 v4.17.0
	github.com/mattn/go-sqlite3 v1.14.15
	github.com/radiochild/repmeta v0.1.2
	github.com/radiochild/utils v0.1.2
	go.uber.org/zap v1.23.0
)

require (
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.9 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.12.23 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.12.19 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.25 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.19 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.26 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.20 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.13.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.11.25 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.13.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.17.1 // indirect
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.13.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.12.0 // indirect
	github.com/jackc/pgx v3.6.2+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/crypto v0.1.0 // indirect
	golang.org/x/text v0.4.0 // indirect
)
