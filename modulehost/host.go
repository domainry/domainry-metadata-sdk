package modulehost

import (
	"context"

	ormmigration "github.com/domainry/domainry-orm/migration"
	"github.com/domainry/domainry-orm/sqlhost"
)

type Database = sqlhost.Database
type SchemaMigration = ormmigration.Migration

type Dialect interface {
	Identifier(string) string
	Table(string) string
	Placeholder(int) string
	Insert(string, []string) string
}

type MigrationRegistrar interface {
	Driver() string
	Schema() string
	ApplyOwnedMigrations(context.Context, string, []SchemaMigration) error
}

type Host interface {
	Database() Database
	Dialect() Dialect
	Migrations() MigrationRegistrar
}
