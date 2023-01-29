package readingFiles

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"projects/ProjectOneAPI_3/models"
)

func GetUsersJSON() models.Users {
	file, err := os.ReadFile("./files/users.json")
	if err != nil {
		fmt.Println("error file can not being reed: ", err)
	}

	var users models.Users

	err = json.Unmarshal(file, &users)
	if err != nil {
		fmt.Println("error json data can not being convert struct: ", err)
	}

	return users
}

func InsertUserJSON(u models.User) (bool, error) {
	users := GetUsersJSON()
	users.User = append(users.User, u)

	jsondata, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error converting struct to JSON: ", err)
		return false, err
	}

	err = os.WriteFile("./files/users.json", jsondata, 0644)
	if err != nil {
		fmt.Println("error writing data to json file: ", err)
		return false, err
	}
	//file, err := os.OpenFile("./files/users.json", os.O_WRONLY, 0644)
	//if err != nil {
	//	fmt.Println("error opening json file: ", err)
	//	return false, err
	//}
	//defer file.Close()
	//
	//_, err = file.Write(jsondata)
	//if err != nil {
	//	fmt.Println("error writing data to json file: ", err)
	//	return false, err
	//}
	return true, nil
}

func UpdateUserJSON(u models.User) (bool, error) {
	users := GetUsersJSON()
	var result bool = false
	for i, user := range users.User {
		if user.Id == u.Id {
			users.User[i].UserName = u.UserName
			users.User[i].EMail = u.EMail
			users.User[i].Password = u.Password
			users.User[i].Database = u.Database
			users.User[i].Collections = u.Collections

			result = true
		}
	}
	if !result {
		err := errors.New("error can not update!")
		return false, err
	}
	jsondata, err := json.Marshal(users)
	if err != nil {
		return false, errors.New("updated value can not being converted to json!")
	}
	err = os.WriteFile("./files/users.json", jsondata, 0644)
	if err != nil {
		fmt.Println("error writing data to json file: ", err)
		return false, err
	}

	return true, nil
}

func DeleteUserJSON(id string) (bool, error) {
	users := GetUsersJSON()
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("error exceptional id: ", err)
	}
	var result bool = false
	for i, user := range users.User {
		if user.Id == _id {
			users.User[i].IsActive = false

			result = true
		}
	}
	if !result {
		err := errors.New("error can not update!")
		return false, err
	}
	jsondata, err := json.Marshal(users)
	if err != nil {
		return false, errors.New("updated value can not being converted to json!")
	}
	err = os.WriteFile("./files/users.json", jsondata, 0644)
	if err != nil {
		fmt.Println("error writing data to json file: ", err)
		return false, err
	}

	return true, nil
}
