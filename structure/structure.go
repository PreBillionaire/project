package structure

import "go.mongodb.org/mongo-driver/bson/primitive"


type Cred struct {
	ID			primitive.ObjectID	`json:"_id,omitempty" bson:"_id,omitempty"`
	Username	string				`json:"username,omitempty" bson:"username,omitempty"`
	Password	string				`json:"password,omitempty" bson:"password,omitempty"`
}

type Employee struct {
	ID			primitive.ObjectID	`json:"_id,omitempty" bson:"_id,omitempty"`
	Username	string				`json:"username,omitempty" bson:"username,omitempty"`
	Firstname	string				`json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname	string				`json:"lastname,omitempty" bson:"lastname,omitempty"`
	Email		string				`json:"email,omitempty" bson:"email,omitempty"`
	Salary		string				`json:"salary,omitempty" bson:"salary,omitempty"`
}
