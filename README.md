# pgfilterproxy

Proxy for PostgreSQL that forwards only configured
[query fingerprints](https://github.com/pganalyze/libpg_query/wiki/Fingerprinting)
and [protocol messages](https://www.postgresql.org/docs/current/protocol-message-formats.html).

Run it was `pgfilterproxy <pgfilterproxy.yaml path>`.
See [pgfilterproxy.yaml.example](pgfilterproxy.yaml.example) for a commented sample config.

You can gather fingerprints of denied queries by looking at the pgfilterproxy log.
