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

type user struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FullName string `json:"name,omitempty" bson:"name,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
	UserName string `json:"username,omitempty" bson:"username,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
	PhoneNumber string `json:"phonenumber,omitempty" bson:"phonenumber,omitempty"`
}

type router struct {
	 database *mongo.Database
	 collection *mongo.Collection
}

func NewUserRouter(database *mongo.Database) *router {
	return &router{
		database: database,
		collection: database.Collection("users"),
	}
}

var users []user

func (router *router) CreateUserHandler(response http.ResponseWriter, request *http.Request) {
	var submittedUser user

	err := json.NewDecoder(request.Body).Decode(&submittedUser)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Invalid data submitted!"))
		return
		//fmt.Println(err)
	}

	result, err := router.collection.InsertOne(context.TODO(), submittedUser)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("An error occurred while creating user"))
		return
	}

	submittedUser.ID = result.InsertedID.(primitive.ObjectID)

	//submittedUser.save()
	json.NewEncoder(response).Encode(submittedUser)
}

func (router *router) ReadUsersHandler(response http.ResponseWriter, request *http.Request) {
	var readUsers []user

	cursor, err := router.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("An error occurred please try again later !"))
		return
	}
	defer cursor.Close(context.TODO())
	err = cursor.All(context.TODO(), &readUsers)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("An error here, please try again later !"))
		return
	}

	json.NewEncoder(response).Encode(readUsers)
}

func (router *router) ReadUserHandler(response http.ResponseWriter, request *http.Request)  {
	var readUser user
	id := chi.URLParam(request, "id")
	userID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Invalid data submitted"))
		return
	}

 	err = router.collection.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&readUser)

 	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("An error here!" + err.Error()))
		return
	}

	json.NewEncoder(response).Encode(readUser)
}

//	if userID > len(users) || userID < 1 {
//		response.WriteHeader(http.StatusBadRequest)
//		response.Write([]byte("Invalid data submitted !"))
//		return
//	}
//

//	userWithID := users[userID-1]

//
func (router *router) UpdateUserHandler(response http.ResponseWriter, request *http.Request) {
	var readUser user
	id := chi.URLParam(request, "id")
	userID, err := primitive.ObjectIDFromHex(id)
	err = json.NewDecoder(request.Body).Decode(&readUser)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Invalid data submitted " + err.Error()))
		return
	}

	//if err != nil {
		//response.WriteHeader(http.StatusBadRequest)
		//response.Write([]byte("Invalid data submitted " + err.Error()))
		//return
	//}

	update := bson.M{
		"$set": readUser,
	}
	_, err = router.collection.UpdateOne(context.TODO(), bson.M{"_id": userID}, update)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("An error here! " + err.Error()))
		return
	}
	json.NewEncoder(response).Encode(readUser)
}

func (router *router) DeleteUserHandler(response http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	userID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Invalid data submitted"))
		return
	}

	result, err := router.collection.DeleteOne(context.TODO(), bson.M{"_id": userID})

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Cannot delete"))
		return
	}

	json.NewEncoder(response).Encode(result)
	response.Write([]byte("Data deleted successfully"))
}
//	id := chi.URLParam(request, "userId")
//	userID, err := strconv.Atoi(id)
//
//	if err != nil {
//		response.WriteHeader(http.StatusBadRequest)
//		response.Write([]byte("Invalid user ID!"))
//		return
//	}
//
//	if userID > len(users) || userID < 1 {
//		response.WriteHeader(http.StatusBadRequest)
//		response.Write([]byte("User Id is out of range !"))
//		return
//	}
//
//	userWithID := &users[userID-1]
//	err = json.NewDecoder(request.Body).Decode(userWithID)
//
//	if err != nil {
//		response.WriteHeader(http.StatusBadRequest)
//		response.Write([]byte("Invalid data submitted !"))
//		return
//	}
//
//	json.NewEncoder(response).Encode(userWithID)
//}

/*
package routers

type post struct {
	Title string `json:"title"`
	Body string `json:"body"`
	Author *user `json:"author"`
}

Create post, Read All Posts, Read Single Post, Update Post, Delete Post

{
	"title": "Test 1",
	"body": "Body 1"
	"author": {
		"id": 5,
		"name": "Joy Nwaiwu",
		"email": "joy@yahoo.com"
    }
}
 */