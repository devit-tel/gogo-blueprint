package company

type Company struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}

func Create(id, name string) *Company {
	return &Company{
		Id:   id,
		Name: name,
	}
}
