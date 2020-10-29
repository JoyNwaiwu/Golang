package routers

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)
//ID int `json:"postid"`
//`json:"name,omitempty" bson:"name,omitempty"`


type post struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string `json:"title,omitempty" bson:"name,omitempty"`
	Body string `json:"body,omitempty" bson:"body,omitempty"`
	Author *User`json:"author,omitempty" bson:"author,omitempty"`


}

func NewPostRouter(database *mongo.Database) *router {
	return &router{
		database: database,
		collection: database.Collection("posts"),
	}
}
var posts []post

func (router *router) CreatePostHandler(response http.ResponseWriter, request *http.Request) {
	var submittedPost post
	//authorID := &users[User]
	//var submittedAuthor post
	//submittedAuthor = post{ User: &User{} }
	//submittedAuthor.ID = *User.ID

	err := json.NewDecoder(request.Body).Decode(&submittedPost)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Invalid data submitted!"))
		return
		//fmt.Println(err)
	}

	result, err := router.collection.InsertOne(context.TODO(), submittedPost)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("An error occurred while creating post"))
		return
	}

	//submittedAuthor.ID = result.InsertedID.(primitive.ObjectID)
	submittedPost.ID = result.InsertedID.(primitive.ObjectID)
	json.NewEncoder(response).Encode(submittedPost)
	//json.NewEncoder(response).Encode(submittedAuthor)
}

func (router *router) ReadPostsHandler(response http.ResponseWriter, request *http.Request) {
	var readPosts []post

	cursor, err := router.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("An error occurred please try again later !"))
		return
	}
	defer cursor.Close(context.TODO())
	err = cursor.All(context.TODO(), &readPosts)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("There is an error here, try again later !"))
		return
	}

	json.NewEncoder(response).Encode(readPosts)
}

func (router *router) ReadPostHandler(response http.ResponseWriter, request *http.Request)  {
	var readPost post
	id := chi.URLParam(request, "id")
	postID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Invalid data submitted"))
		return
	}

	err = router.collection.FindOne(context.TODO(), bson.M{"_id": postID}).Decode(&readPost)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("An error here!" + err.Error()))
		return
	}

	json.NewEncoder(response).Encode(readPost)
}

func (router *router) UpdatePostHandler(response http.ResponseWriter, request *http.Request) {
	var readPost post
	id := chi.URLParam(request, "id")
	postID, err := primitive.ObjectIDFromHex(id)
	err = json.NewDecoder(request.Body).Decode(&readPost)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Invalid data submitted " + err.Error()))
		return
	}

	update := bson.M{
		"$set": readPost,
	}
	_, err = router.collection.UpdateOne(context.TODO(), bson.M{"_id": postID}, update)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("An error here! " + err.Error()))
		return
	}
	json.NewEncoder(response).Encode(readPost)
}

func (router *router) DeletePostHandler(response http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	postID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Invalid data submitted"))
		return
	}

	result, err := router.collection.DeleteOne(context.TODO(), bson.M{"_id": postID})

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Cannot delete"))
		return
	}

	json.NewEncoder(response).Encode(result)
	response.Write([]byte("Data deleted successfully"))
}
//Create post, Read All Posts, Read Single Post, Update Post, Delete Post
