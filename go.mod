module github.com/debfx/pgfilterproxy

go 1.15

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/pganalyze/pg_query_go/v2 v2.1.2
	github.com/rueian/pgbroker v0.0.17
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/rueian/pgbroker => ./pgbroker
