package models

//Relacion entre usuarios
type Relation struct {
	UserId         string `bson:"userId" json:"userId"`
	UserRelationId string `bson:"userRelationId" json:"userRelationId"`
}
