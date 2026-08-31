# Domainry Metadata SDK

Deployment-neutral contracts between Domainry Runtime and the Metadata owner.

## Package layout

- The root package is the stable Metadata `Binding` entrypoint.
- `persistence` owns definition snapshots and transaction-aware, source-owned Metadata DML contracts.
- `modulehost` describes the database, dialect, and host migration registrar borrowed by an embedded Metadata module.

Metadata owns its table names and DML while participating in the host transaction through bounded executor interfaces exposed by `persistence`. Every database continues to use the host-owned `_schema_migrations` ledger.

Run `go test ./...` before publishing an immutable SDK version.
