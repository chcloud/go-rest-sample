package mapper

type Mapper interface {
	FindOne(ID int64) interface{}
	Save(interface{}) error
	Delete(ID int64) error
}
