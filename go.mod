module github.com/debfx/pgfilterproxy

go 1.15

require (
	github.com/lfittl/pg_query_go v1.0.1
	github.com/rueian/pgbroker v0.0.15
	gopkg.in/yaml.v2 v2.3.0
)

replace (
	github.com/rueian/pgbroker => ./pgbroker
)
