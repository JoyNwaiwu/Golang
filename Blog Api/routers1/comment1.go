package routers

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type comment struct {
	Body string `json:"body,omitempty" bson:"body,omitempty"`
	Article *post
	//CreatedAt time.Now()
}

func CommentRouter(database *mongo.Database) *router {
	return &router{
		database: database,
		collection: database.Collection("comments"),
	}
}
var comments []comment

func (router *router) CreateCommentHandler(response http.ResponseWriter, request *http.Request) {
	var subComment comment

	err := json.NewDecoder(request.Body).Decode(&subComment)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Invalid data submitted!"))
		return
	}

	_, error := router.collection.InsertOne(context.TODO(), subComment)
	if error != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("An error occurred while creating user"))
		return
	}

	json.NewEncoder(response).Encode(subComment)
}

func (router *router) ReadCommentsHandler(response http.ResponseWriter, request *http.Request) {
	var readComments []comment

	cursor, err := router.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("An error occurred please try again later !"))
		return
	}
	defer cursor.Close(context.TODO())
	err = cursor.All(context.TODO(), &readComments)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("There is an error here, try again later !"))
		return
	}

	json.NewEncoder(response).Encode(readComments)
}

