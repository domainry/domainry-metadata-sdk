# Domainry Metadata SDK

Deployment-neutral contracts between Domainry Runtime and the Metadata owner.

## Package layout

- The root package is the stable Metadata business contract and `Binding`
  entrypoint. It exposes definition, localization, dictionary and authorized
  projection ports without exposing HTTP handlers or implementation services.
- `modulehost` describes the database, dialect, and host migration registrar borrowed by an embedded Metadata module.

Metadata owns its table names, transactions and DML. Every database continues
to use the host-owned `_schema_migrations` ledger.

Run `go test ./...` before publishing an immutable SDK version.
