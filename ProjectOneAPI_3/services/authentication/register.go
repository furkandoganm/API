package authentication

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
	"projects/ProjectOneAPI_3/configs"
	"projects/ProjectOneAPI_3/models"
	"projects/ProjectOneAPI_3/readingFiles"
	"projects/ProjectOneAPI_3/services/dbClient"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	body, err := io.ReadAll(r.Body)
	Err(w, err, "error reading body:")

	err = json.Unmarshal(body, &newUser)
	Err(w, err, "error unmarshal to model:")

	dbClient := dbClient.Client("users", "user_identify")
	model, err := dbClient.InsertUser(newUser)
	Err(w, err, "error unmarshal to model:")

	_, err = readingFiles.InsertUserJSON(model)
	Err(w, err, "error inserting json file:")

	var nData models.Movie
	nData.Id = primitive.NilObjectID
	nData.DirectorId = primitive.NilObjectID
	nData.IMDBScore = 0
	nData.ScriptwriterId = primitive.NilObjectID
	nData.MovieName = "first movie"

	for _, c := range newUser.Collections {
		configs.CConnectionMDB(nData, newUser.Database, c)
	}

	fmt.Fprintln(w, "success.")

}

//func Register(c *fiber.Ctx) error {
//	var newUser models.User
//	if err := c.BodyParser(&newUser); err != nil {
//		return c.Status(http.StatusBadRequest).JSON(err.Error())
//	}
//
//	dbC := dbClient.UserClient()
//	result, err := dbC.InsertUser(newUser)
//	if err != nil {
//		return c.Status(http.StatusBadRequest).JSON(err.Error())
//	}
//
//	return c.Status(http.StatusOK).JSON(result)
//}
