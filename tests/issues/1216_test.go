package issues

import (
	"context"
	"testing"

	"github.com/ClickHouse/clickhouse-go/v2"
	clickhouse_tests "github.com/ClickHouse/clickhouse-go/v2/tests"
	"github.com/stretchr/testify/require"
)

func Test1216(t *testing.T) {
	var (
		conn, err = clickhouse_tests.GetConnectionTCP("issues", clickhouse.Settings{
			"max_execution_time": 60,
		}, nil, &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		})
	)
	ctx := context.Background()
	require.NoError(t, err)
	const ddl = "CREATE TABLE IF NOT EXISTS test_1216 (`@id` String,`\"@id_with_quotes\"` String) Engine = Memory"
	require.NoError(t, conn.Exec(ctx, ddl))
	defer func() {
		conn.Exec(ctx, "DROP TABLE IF EXISTS test_1216")
	}()

	_, err = conn.PrepareBatch(context.Background(), "INSERT INTO test_1216 (`@id`, `\"@id_with_quotes\"`)")
	require.NoError(t, err)
}
