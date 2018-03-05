package mapper

import (
	"testing"

	_ "github.com/chcloud/go-rest-sample/pkg/config"
	_ "github.com/go-sql-driver/mysql"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	Connect()
	info, err := engine.DBMetas()
	assert.NoError(t, err)
	assert.NotNil(t, info)
}
