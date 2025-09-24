module github.com/debfx/pgfilterproxy

go 1.23.0

require (
	github.com/pganalyze/pg_query_go/v6 v6.1.0
	github.com/rueian/pgbroker v0.0.18
	gopkg.in/yaml.v2 v2.4.0
)

require google.golang.org/protobuf v1.33.0 // indirect

replace github.com/rueian/pgbroker => ./pgbroker
