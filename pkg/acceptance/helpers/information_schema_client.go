package helpers

import (
	"context"
	"fmt"
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/internal/collections"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
	"github.com/stretchr/testify/require"
)

type InformationSchemaClient struct {
	context *TestClientContext
	ids     *IdsGenerator
}

func NewInformationSchemaClient(context *TestClientContext, idsGenerator *IdsGenerator) *InformationSchemaClient {
	return &InformationSchemaClient{
		context: context,
		ids:     idsGenerator,
	}
}

func (c *InformationSchemaClient) client() *sdk.Client {
	return c.context.client
}

type QueryHistory struct {
	QueryId   string
	QueryText string
	QueryTag  string
}

func (c *InformationSchemaClient) GetQueryHistory(t *testing.T, limit int) []QueryHistory {
	t.Helper()

	result, err := c.client().QueryUnsafe(context.Background(), fmt.Sprintf("SELECT * FROM TABLE(INFORMATION_SCHEMA.QUERY_HISTORY(RESULT_LIMIT => %d))", limit))
	require.NoError(t, err)

	return collections.Map(result, func(queryResult map[string]*any) QueryHistory {
		return c.mapQueryHistory(t, queryResult)
	})
}

func (c *InformationSchemaClient) GetQueryHistoryByQueryId(t *testing.T, limit int, queryId string) QueryHistory {
	t.Helper()

	result, err := c.client().QueryUnsafe(context.Background(), fmt.Sprintf("SELECT * FROM TABLE(INFORMATION_SCHEMA.QUERY_HISTORY(RESULT_LIMIT => %d)) WHERE QUERY_ID = '%s'", limit, queryId))
	require.NoError(t, err)
	require.Len(t, result, 1)

	return c.mapQueryHistory(t, result[0])
}

func (c *InformationSchemaClient) mapQueryHistory(t *testing.T, queryResult map[string]*any) QueryHistory {
	t.Helper()
	var queryHistory QueryHistory
	if v, ok := queryResult["QUERY_ID"]; ok && v != nil {
		queryHistory.QueryId = (*v).(string)
	}
	if v, ok := queryResult["QUERY_TEXT"]; ok && v != nil {
		queryHistory.QueryText = (*v).(string)
	}
	if v, ok := queryResult["QUERY_TAG"]; ok && v != nil {
		queryHistory.QueryTag = (*v).(string)
	}
	return queryHistory
}
