package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//type Movies struct {
//	Movie []Movie `json:"movie"`
//}

type Movie struct {
	Id        primitive.ObjectID `json:"id"`
	MovieName string             `json:"moviename"`
	//Director     InfoUser           `json:"director"`
	//Scriptwriter InfoUser           `json:"scriptwriter"`
	DirectorId     primitive.ObjectID `json:"directorid"`
	ScriptwriterId primitive.ObjectID `json:"scriptwriterid"`
	IMDBScore      float32            `json:"imdbscore"`
	//ProducedYear   time.Time          `json:"production_year"`
}

type Music struct {
	Id         primitive.ObjectID `json:"id"`
	Name       string             `json:"name"`
	SingerId   primitive.ObjectID `json:"singer_id"`
	ComposerId primitive.ObjectID `json:"composer_id"`
}
