package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	User []User `json:"users"`
}

type User struct {
	Id          primitive.ObjectID `json:"id"`
	UserName    string             `json:"username"`
	EMail       string             `json:"email"`
	Password    string             `json:"password"`
	Status      string             `json:"status"`
	IsActive    bool               `json:"isactive"`
	Database    string             `json:"database"`
	Collections []string           `json:"collections"`
}

type InfoUsers struct {
	InfoUser []InfoUser `json:"info_user"`
}

type InfoUser struct {
	Id            primitive.ObjectID `json:"id"`
	Name          string             `json:"name"`
	Surname       string             `json:"surname"`
	YearOfBirth   int                `json:"year_of_birth"`
	MaritalStatus bool               `json:"marital_status"`
	Citizen       string             `json:"citizen"`
}

const StandardStatus = "user"
