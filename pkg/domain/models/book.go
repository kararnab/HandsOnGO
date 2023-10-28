package models

import "github.com/google/uuid"

type Book struct {
	ID       uuid.UUID `json:"id" bson:"id,omitempty"`
	ISBN     string    `json:"isbn" bson:"isbn,omitempty"`
	BookName string    `json:"name" bson:"name,omitempty"`
	Price    float64   `json:"price" bson:"price,omitempty"`
}
