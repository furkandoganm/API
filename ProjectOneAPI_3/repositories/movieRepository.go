package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"projects/ProjectOneAPI_3/models"
	"time"
)

func (dbClient *Client) InsertMovie(req models.Movie) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	req.Id = primitive.NewObjectID()

	result, err := dbClient.Collection.InsertOne(ctx, req)
	if err != nil || result.InsertedID == nil {
		return err
	}
	return nil
}

func (dbClient *Client) GetMovies() ([]models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	result, err := dbClient.Collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("error getting users: ", err)
		return nil, err
	}

	var req models.Movie
	var reqs []models.Movie

	for result.Next(ctx) {
		if err = result.Decode(&req); err != nil {
			fmt.Println("error decoding user that brought: ", err)
			return nil, err
		}
		reqs = append(reqs, req)
	}
	return reqs, nil
}

func (dbClient *Client) UpdateMovie(req models.Movie) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	var updatedMovie = bson.M{"moviename": req.MovieName, "imdbscore": req.IMDBScore, "directorid": req.DirectorId, "scriptwriterid": req.ScriptwriterId}
	matchCount, err := dbClient.Collection.UpdateOne(ctx, bson.M{"id": req.Id}, bson.M{"$set": updatedMovie})
	if err != nil || matchCount == nil {
		fmt.Println("error updating movie: ", err)
		return false, err
	}
	return true, nil
}

func (dbClient *Client) DeleteMovie(id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("error id can not converted: ", err)
	}

	_, err = dbClient.Collection.DeleteOne(ctx, bson.M{"id": _id})

	if err != nil {
		fmt.Println("error updating user: ", err)
		return false, err
	}
	return true, nil
}
