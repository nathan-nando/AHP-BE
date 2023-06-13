package model

type Collection struct {
	Name        string
	Description string
}

type CollectionModel struct {
	BaseModel
	Collection
}
