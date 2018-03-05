package mapper

import (
	"time"

	"gopkg.in/logger.v1"
)

//User user struct
type User struct {
	ID       int64     `xorm:"pk autoincr 'id'"`
	Name     string    `xorm:"char(50) not null 'user_name'"`
	CreateAt time.Time `xorm:"'created'"`
}

// UserMapper map User struct with database user table
type UserMapper struct {
}

// Save insert user to table
func (p *UserMapper) Save(u *User) error {
	re, err := engine.InsertOne(u)
	if err != nil {
		log.Error("An error occurred while saving user to db ", err)
		return err
	}
	log.Info(re)
	return nil
}

// FindOne  find user by id
func (p *UserMapper) FindOne(id int64) *User {
	var user User
	log.Info("get user id: ", id)
	// engine.ID(id).Get(user)
	ok, err := engine.Alias("u").Where("u.id=?", id).Get(&user)
	if err != nil {
		log.Errorf("Get user id=%d error: %s", id, err)
	}
	log.Infof("Get user result: %v", ok)
	if ok {
		return &user
	}
	return nil
}
