package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"projects/ProjectOneAPI_3/models"
	"time"
)

type Client struct {
	Collection *mongo.Collection
}

func (dbClient *Client) InsertUser(req models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	req.Id = primitive.NewObjectID()
	req.IsActive = true
	req.Status = models.StandardStatus

	result, err := dbClient.Collection.InsertOne(ctx, req)
	if err != nil || result.InsertedID == nil {
		//fmt.Println("error inserting user: ", err)
		return req, err
	}
	return req, nil
}

func (dbClient *Client) GetUsers() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	result, err := dbClient.Collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("error getting users: ", err)
		return nil, err
	}

	var req models.User
	var reqs []models.User

	for result.Next(ctx) {
		if err = result.Decode(&req); err != nil {
			fmt.Println("error decoding user that brought: ", err)
			return nil, err
		}
		if req.IsActive {
			reqs = append(reqs, req)
		}
	}
	return reqs, nil
}

//func (dbClient *UserClient) GetUser() (models.User, error) {
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
//	defer cancel()
//
//	result, err := dbClient.Collection.FindOne(ctx, bson.M{"id": })
//	if err != nil {
//		fmt.Println("error getting users: ", err)
//		return nil, err
//	}
//
//	var req models.User
//	var reqs []models.User
//
//	for result.Next(ctx) {
//		if err = result.Decode(&req); err != nil {
//			fmt.Println("error decoding user that brought: ", err)
//			return nil, err
//		}
//		if req.IsActive {
//			reqs = append(reqs, req)
//		}
//	}
//	return reqs, nil
//}

func (dbClient *Client) UpdateUser(req models.User) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	var updatedUser = bson.M{"username": req.UserName, "email": req.EMail, "password": req.Password, "database": req.Database, "collections": req.Collections}
	matchCount, err := dbClient.Collection.UpdateOne(ctx, bson.M{"id": req.Id}, bson.M{"$set": updatedUser})
	if err != nil || matchCount == nil {
		fmt.Println("error updating user: ", err)
		return false, err
	}
	return true, nil
}

func (dbClient *Client) DeleteUser(id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("error id can not converted: ", err)
	}

	var deletedUser = bson.M{"isactive": false}
	matchCount, err := dbClient.Collection.UpdateOne(ctx, bson.M{"id": _id}, bson.M{"$set": deletedUser})
	if err != nil || matchCount == nil {
		fmt.Println("error updating user: ", err)
		return false, err
	}
	return true, nil
}

//func (dbClient *UserClient) InsertInfoUser(req models.InfoUser) (bool, error) {
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
//	defer cancel()
//
//	req.Id = primitive.NewObjectID()
//
//	result, err := dbClient.Collection.InsertOne(ctx, req)
//	if err != nil || result.InsertedID == nil {
//		//fmt.Println("error inserting user: ", err)
//		return false, err
//	}
//	return true, nil
//}
//
//func (dbClient *UserClient) InsertMovies(req models.Movie) (bool, error) {
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
//	defer cancel()
//
//	req.Id = primitive.NewObjectID()
//
//	result, err := dbClient.Collection.InsertOne(ctx, req)
//	if err != nil || result.InsertedID == nil {
//		//fmt.Println("error inserting user: ", err)
//		return false, err
//	}
//	return true, nil
//}
