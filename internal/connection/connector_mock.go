package connection

import (
	"testing"

	"github.com/bookmark-project-learn/user-service/internal/test/data/fixture"
)

func InitDBConnectorMock(t *testing.T, fix fixture.Fixture) (DBConnector, error) {
	// fix := fixture.NewUserTestCase(t)
	gormSql := fixture.NewFixture(t, fix)
	return NewDBConnector(gormSql), nil
}
