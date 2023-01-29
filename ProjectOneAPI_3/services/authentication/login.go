package authentication

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"projects/ProjectOneAPI_3/models"
	"projects/ProjectOneAPI_3/readingFiles"
	"projects/ProjectOneAPI_3/services/dbClient"
)

func LogIn(w http.ResponseWriter, r *http.Request) {
	var user models.User
	body, err := io.ReadAll(r.Body)
	Err(w, err, "error reading body:")

	err = json.Unmarshal(body, &user)
	Err(w, err, "error unmarshalling body:")

	users := readingFiles.GetUsersJSON().User
	var result bool = false
	for _, u := range users {
		if u.UserName == user.UserName || u.Password == user.Password {
			if u.Database == user.Database {
				for _, c := range user.Collections {
					if u.Collections[0] == c {
						result = true
					}
				}
			} else {
				Err(w, nil, "There is no database that you have written!")
			}
		}
	}
	if result {
		repo := dbClient.Client(user.Database, user.Collections[0])
		movies, err := repo.GetMovies()
		Err(w, err, "error pulling data from movies:")
		fmt.Fprintln(w, movies)
	} else {
		Err(w, nil, "There is no collection that you have written!")
	}

}

func Err(w http.ResponseWriter, err error, str string) {
	if err != nil {
		er := fmt.Sprintf("%s %v", str, err)
		fmt.Println(er)
		fmt.Fprintln(w, er)
	}
}
