package mapper

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/chcloud/go-rest-sample/pkg/config"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	Connect()
	info, err := engine.DBMetas()
	assert.NoError(t, err)
	assert.NotNil(t, info)
}
