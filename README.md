# pgfilterproxy

Proxy for PostgreSQL proxy that allows for filtering by query
[fingerprint](https://github.com/pganalyze/libpg_query/wiki/Fingerprinting)
and protocol messages.

Run it was `pgfilterproxy <pgfilterproxy.yaml path>`.
See [pgfilterproxy.yaml.example](pgfilterproxy.yaml.example) for a commented samepl config.

You can gather fingerprints of denied queries by looking at the pgfilterproxy log.
