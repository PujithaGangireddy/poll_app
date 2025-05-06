package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"poll_app/ent"
	"poll_app/ent/user"
	"poll_app/internal/utils"

	"github.com/julienschmidt/httprouter"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUp(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var req AuthRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Println("no issue with auth service")

		hashedPassword, _ := utils.HashedPassword(req.Password)

		u, err := client.User.
			Create().
			SetEmail(req.Email).
			SetPassword(hashedPassword).
			Save(context.Background())

		if err != nil {
			http.Error(w, "Error Creating User", http.StatusInternalServerError)
		}

		json.NewEncoder(w).Encode(u)
	}
}

func Login(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var req AuthRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := client.User.
			Query().
			Where(user.Email(req.Email)).
			Only(context.Background())

		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		if !utils.CheckPasswordHash(req.Password, user.Password) {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		token, err := utils.GenerateJWT(user.ID)
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"token": token})
	}
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := map[string]string{
		"status":  "ok",
		"env":     "dev",
		"version": "v1",
	}

	json.NewEncoder(w).Encode(data)

}
