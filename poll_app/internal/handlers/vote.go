package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"poll_app/ent"
	"poll_app/ent/polloption"
	"poll_app/ent/user"
	"poll_app/ent/vote"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type VoteRequest struct {
	UserID    int   `json:"user_id"`
	OptionIDs []int `json:"option_ids"`
}

// CastVote allows a user to cast a vote for specific options.
func CastVote(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var req VoteRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		vote, err := client.Vote.
			Create().
			SetUserID(req.UserID).
			AddOptionIDs(req.OptionIDs...).
			Save(context.Background())

		if err != nil {
			http.Error(w, "Error casting vote", http.StatusInternalServerError)
			return
		}
		options, err := client.PollOption.
			Query().
			Where(polloption.IDIn(req.OptionIDs...)).
			WithPoll().
			All(context.Background())
		if err != nil {
			http.Error(w, "Error fetching option details", http.StatusInternalServerError)
			return
		}

		optionTexts := []string{}
		var pollQuestion string
		if len(options) > 0 {
			pollQuestion = options[0].Edges.Poll.Question // assuming all options belong to the same poll
		}
		for _, option := range options {
			optionTexts = append(optionTexts, option.Text)
		}

		response := map[string]interface{}{
			"message":       "Vote cast successfully",
			"vote_id":       vote.ID,
			"user_email":    user.Email,
			"poll_question": pollQuestion,
			"option_texts":  optionTexts,
		}

		json.NewEncoder(w).Encode(response)
	}
}

// ListVotes fetches all votes with the associated user and options.
func ListVotes(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		votes, err := client.Vote.Query().WithUser().WithOptions().All(context.Background())
		if err != nil {
			http.Error(w, "Error fetching votes", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(votes)
	}
}

// GetVote fetches a specific vote by ID.
func GetVote(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		voteID, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, "Invalid vote ID", http.StatusBadRequest)
			return
		}

		vote, err := client.Vote.Query().
			Where(vote.ID(voteID)).
			WithUser().
			WithOptions().
			Only(context.Background())
		if err != nil {
			http.Error(w, "Error fetching vote", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(vote)
	}
}
