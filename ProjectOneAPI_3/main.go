package main

import (
	"net/http"
	"projects/ProjectOneAPI_3/services/authentication"
)

func main() {

	//var t models.DBUser
	//t.User = "trial1"
	//t.DB = "trial1"
	//t.PasswordDigestor = "server"
	//t.Roles = []models.Role{
	//	{
	//		Role: "readWriteAnyDatabase",
	//		DB:   "Admin",
	//	},
	//}
	//t.Pwd = "Test123!"
	//configs.CConnectionMDB(t, "test2", "test2")

	//route := fiber.New()
	//route.Post("/register", authentication.Register)
	//route.Listen(":8080")

	mux := http.NewServeMux()
	mux.HandleFunc("/register", authentication.Register)
	mux.HandleFunc("/login", authentication.LogIn)

	http.ListenAndServe(":8080", mux)

	//dbClient := configs.ConnectionMDB("users", "user_identify")
	//
	//userClient := repositories.UserClient{Collection: dbClient}

	//var t models.User
	//t.UserName = "jaskier"
	//t.EMail = "jaskier@gmail.com"
	//t.Password = "12345.!"
	//
	//userClient.InsertUser(t)

	//for _, m := range readingFiles.GetUsersJSON().User {
	//	userClient.InsertUser(m)
	//}

	//var t models.InfoUser
	//t.Name = "Stanislaw"
	//t.Surname = "Lem"
	//t.Citizen = "Ukraine"
	//t.MaritalStatus = false
	//t.YearOfBirth = 1921
	//
	//userClient.InsertInfoUser(t)

	//id, err := primitive.ObjectIDFromHex("633c374340a7610e518b3284")
	//if err != nil {
	//	fmt.Println("error id can not being converted to primitiveId: ", err)
	//}
	//var tempModel = models.User{
	//	Id:       id,
	//	UserName: "vVGoh",
	//	EMail:    "stary@night.com",
	//	Password: "12345.!_2",
	//}

	//id, err := primitive.ObjectIDFromHex("633c371130c9ced71a083d2a")
	//if err != nil {
	//	fmt.Println("error id can not being converted to primitiveId: ", err)
	//}
	//
	//var coll = []string{
	//	"education",
	//}
	//
	//var tempModel = models.User{
	//	Id:          id,
	//	UserName:    "fdogan_3",
	//	EMail:       "fdogangazi@gmail.com",
	//	Password:    "12345.!_2",
	//	Database:    "muaddib",
	//	Collections: coll,
	//}

	//result, err := readingFiles.DeleteUserJSON("633c371130c9ced71a083d2b")
	//if !result {
	//	fmt.Println(err)
	//}
	//fmt.Println("success")

	//
	//var users []models.User = readingFiles.GetUsersJSON()
	//
	//for user, index := range users {
	//	fmt.Printf("%v: %v\n", index, user)
	//}
	//
	//tempUser := models.User{
	//	UserName:   "fdogan",
	//	EMail:      "fdogangazi@gmail.com",
	//	Status:     "Admine",
	//	Password:   "12345.!",
	//	Database:   "muaddib",
	//	Collection: "education",
	//}
	//tempUser := models.User{
	//	UserName: "vVGoh",
	//	EMail:    "stary@night.com",
	//	Password: "12345.!",
	//}
	//
	//var coll []string
	//coll = append(coll, "director", "scripter")
	//tempUser := models.User{
	//	UserName:    "CKaufman",
	//	EMail:       "charly@kaufman.com",
	//	Password:    "itotet",
	//	Database:    "product",
	//	Collections: coll,
	//}
	//result, err := userClient.InsertUser(tempUser)
	//if result {
	//	fmt.Println("Success")
	//} else {
	//	fmt.Println(err)
	//}
	//
	//result, err := userClient.GetUsers()
	//if err == nil {
	//	fmt.Println("Success")
	//	fmt.Println(result)
	//} else {
	//	fmt.Println(err)
	//}
	//
	//id, _ := primitive.ObjectIDFromHex("633c374340a7610e518b3284")
	//tempUser := models.User{
	//	UserName: "vVGoh_2",
	//	EMail:    "stary@night_2.com",
	//	Password: "12345.!_2",
	//	Id:       id,
	//}
	//userClient.UpdateUser(tempUser)
	//
	//result, err := userClient.DeleteUser("633d5318b30727e55f7c48df")
	//if !result {
	//	fmt.Println("error deleting user: ", err)
	//} else {
	//	fmt.Println("success")
	//}

}
