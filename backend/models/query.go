package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Query struct {
	Id      primitive.ObjectID `bson:"_id"`
	QueryId string             `json:"query_id"`
	Name    string             `json:"name" validate:"required"`
	Email   string             `json:"email" validate:"required,email"`
	Subject string             `json:"subject" validate:"required"`
	Message string             `json:"message" validate:"required"`
}
