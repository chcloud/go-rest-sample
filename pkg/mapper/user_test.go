package mapper

import (
	"testing"
	"time"
)

func init() {
	connectDb() //connect DB
}
func TestUserMapper_Save(t *testing.T) {
	user := &User{
		Name:     "test",
		CreateAt: time.Now(),
	}
	uMapper := &UserMapper{}
	uMapper.Save(user)
}
