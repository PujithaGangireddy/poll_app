package main

import (
	"log"
	"net/http"
	"poll_app/internal/db"
	"poll_app/internal/handlers"

	"github.com/julienschmidt/httprouter"
)

func main() {
	client, err := db.NewClient()

	if err != nil {
		log.Fatalf("Error connecting db: %v", err)
	}

	defer client.Close()

	log.Println("Database connected and schema created successfully!")

	router := httprouter.New()
	
	router.GET("/health", handlers.HealthCheckHandler)

	router.POST("/signup", handlers.SignUp(client))
	router.GET("/users", handlers.ListUsers(client))
	router.GET("/users/:id", handlers.GetUser(client))
	router.DELETE("/users/:id", handlers.DeleteUser(client))

	router.POST("/login", handlers.Login(client))

	router.POST("/polls", handlers.CreatePoll(client))
	router.GET("/polls", handlers.ListPolls(client))
	router.DELETE("/polls/:id", handlers.DeletePoll(client))

	router.POST("/vote", handlers.CastVote(client))
	router.GET("/votes", handlers.ListVotes(client))
	router.GET("/vote/:id", handlers.GetVote(client))

	handlerWithCORS := handlers.EnableCORS(router)

	log.Fatal(http.ListenAndServe(":8080", handlerWithCORS))
}
