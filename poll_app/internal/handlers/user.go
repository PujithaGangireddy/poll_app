package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"poll_app/ent"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func ListUsers(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		users, err := client.User.Query().All(r.Context())
		if err != nil {
			http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

func GetUser(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id, _ := strconv.Atoi(ps.ByName("id"))
		user, err := client.User.Get(r.Context(), id)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(user)
	}
}

func DeleteUser(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id, _ := strconv.Atoi(ps.ByName("id"))
		err := client.User.DeleteOneID(id).Exec(context.Background())
		// log.Println("Error deleting user:", err)
		if err != nil {
			http.Error(w, "Could not delete user", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // 200 OK status
		json.NewEncoder(w).Encode(map[string]string{
			"message": fmt.Sprintf("User with ID %d successfully deleted", id),
		})
	}
}
