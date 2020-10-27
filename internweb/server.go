package main

import (
	"context"
	"github.com/JoyNwaiwu/internweb/routers"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

var client *mongo.Client

func main () {
	router := chi.NewRouter()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database("testing")
	userRouter := routers.NewUserRouter(database)

	router.Post("/users", userRouter.CreateUserHandler)
	router.Get("/users", userRouter.ReadUsersHandler)
	router.Get("/users/{id}", userRouter.ReadUserHandler)
	router.Put("/users/{id}", userRouter.UpdateUserHandler)
	router.Delete("/users/{id}", userRouter.DeleteUserHandler)

	router.Post("/posts", routers.CreatePost)
	router.Get("/posts", routers.ReadPosts)
	router.Get("/posts/{postId}", routers.ReadPost)
	router.Put("/posts/{postId}", routers.UpdatePost)
	router.Delete("/posts/{postId}", routers.DeletePost)

	http.ListenAndServe(":8000", router)

	//http.HandleFunc("/users/new", routers.CreateUserHandler)
	//http.HandleFunc("/users/all", routers.ReadUsersHandler)

}


