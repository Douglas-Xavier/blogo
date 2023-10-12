package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"yakki/blogo/data"
	"yakki/blogo/models"
)

type UserInterface interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	// GetUser(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	RouteByMethod(w http.ResponseWriter, r *http.Request)
}

type UserEndpoints struct {
	r data.UserRepositoryInterface
}

func (u *UserEndpoints) RouteByMethod(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		u.CreateUser(w, r)
	} else if r.Method == http.MethodGet {
		u.GetUsers(w, r)
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (u *UserEndpoints) GetUsers(w http.ResponseWriter, r *http.Request) {

	data, err := u.r.GetUsers()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(data)
}

func (u *UserEndpoints) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var user models.NewUser

	err := json.Unmarshal(body, &user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u.r.CreateUser(user)

}

func ProvideUserApi(r data.UserRepositoryInterface) UserInterface {
	return &UserEndpoints{r}
}
