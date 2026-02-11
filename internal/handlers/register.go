package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
)

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type Id struct {
	Id int64 `json:"id"`
}

func (srv *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		log.Println(err)
		return
	}
	idInt, err := srv.deps.UserService.Register(context.Background(), user.Name, user.Username, user.Password)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var id Id
	id.Id = idInt
	if err := json.NewEncoder(w).Encode(id); err != nil {
		log.Println(err)
		return
	}
}
