package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" bson:"id,omitempty"`
	Username string    `json:"username" bson:"username,omitempty"`
	Password string    `json:"password" bson:"password,omitempty"`
	Role     UserRole  `json:"role" bson:"gender,omitempty"`
	Name     string    `json:"name" bson:"name,omitempty"`
	Age      int       `json:"age" bson:"age,omitempty"`
	Gender   Gender    `json:"gender" bson:"gender,omitempty"`
	Contact  string    `json:"contact" bson:"contact,omitempty"`
}
