package resource

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserResource_WebService(t *testing.T) {
	urs := &UserResource{}
	assert.NotNil(t, urs.WebService())
}
