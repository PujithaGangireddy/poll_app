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

type CreatePollRequest struct {
	Question string   `json:"question"`
	Options  []string `json:"options"`
}

func CreatePoll(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var req CreatePollRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		p, err := client.Poll.
			Create().
			SetQuestion(req.Question).
			Save(context.Background())

		if err != nil {
			http.Error(w, "Error Creating Poll", http.StatusInternalServerError)
			return
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, option := range req.Options {
			client.PollOption.Create().
				SetText(option).
				SetPoll(p).
				Save(context.Background())
		}

		json.NewEncoder(w).Encode(p)
	}
}

// func ListPolls(client *ent.Client) httprouter.Handle {
// 	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 		polls, err := client.Poll.Query().All(context.Background())
// 		if err != nil {
// 			http.Error(w, "Error fetching polls", http.StatusInternalServerError)
// 			return
// 		}
// 		json.NewEncoder(w).Encode(polls)
// 	}
// }

func ListPolls(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		polls, err := client.Poll.
			Query().
			WithOptions(). // Include related options in the query
			All(context.Background())
		if err != nil {
			http.Error(w, "Error fetching polls", http.StatusInternalServerError)
			return
		}

		// Encode the polls with their options as JSON
		json.NewEncoder(w).Encode(polls)
	}
}

func DeletePoll(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id, _ := strconv.Atoi(ps.ByName("id"))
		if err := client.Poll.DeleteOneID(id).Exec(context.Background()); err != nil {
			http.Error(w, "Could not delete poll", http.StatusInternalServerError)
			return
		}
		// w.WriteHeader(http.StatusNoContent)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // 200 OK status
		json.NewEncoder(w).Encode(map[string]string{
			"message": fmt.Sprintf("Poll with ID %d successfully deleted", id),
		})
	}
}
