package modulehost

import (
	"context"
	"database/sql"
	"testing"
)

func TestExecutorContextPrefersHostTransaction(t *testing.T) {
	database := &sql.DB{}
	transaction := &sql.Tx{}
	if got := ExecutorFromContext(context.Background(), database); got != database {
		t.Fatal("fallback database was not returned")
	}
	ctx := WithExecutor(context.Background(), transaction)
	if got := ExecutorFromContext(ctx, database); got != transaction {
		t.Fatal("host transaction was not returned")
	}
}
