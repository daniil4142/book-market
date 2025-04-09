module github.com/daniil4142/book-market/category-service

go 1.23.5

require (
	github.com/daniil4142/book-market/category-service/pkg/category-service v0.0.0-20250402184720-5a5f56b66669
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.26.3
	github.com/jackc/pgx/v4 v4.13.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.24.0
	google.golang.org/grpc v1.71.1
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/mfridman/interpolate v0.0.2 // indirect
	github.com/sethvargo/go-retry v0.3.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/sync v0.12.0 // indirect
)

require (
	github.com/Masterminds/squirrel v1.5.4
	github.com/envoyproxy/protoc-gen-validate v1.2.1 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.10.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.1.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgtype v1.8.1 // indirect
	github.com/jmoiron/sqlx v1.4.0
	github.com/kr/text v0.2.0 // indirect
	github.com/pressly/goose/v3 v3.24.2
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto v0.0.0-20211021150943-2b146023228c // indirect
	google.golang.org/protobuf v1.36.6
)

//replace github.com/daniil4142/book-market/category-service/pkg/category-service/pkg/category-service => ./pkg/category-service
