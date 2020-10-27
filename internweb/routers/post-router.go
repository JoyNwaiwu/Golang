package routers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type post struct {
	ID int `json:"postid"`
	Title string `json:"title"`
	Body string `json:"body"`
	Author *user `json:"author"`
}

var posts []post

func (pos *post) save() {
	pos.ID = len(posts) + 1
	posts = append(posts, *pos)
}

func CreatePost(res http.ResponseWriter, req *http.Request) {
	var submittedPost post

	err := json.NewDecoder(req.Body).Decode(&submittedPost)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Invalid post submitted!"))
		return
	}

	submittedPost.save()
	json.NewEncoder(res).Encode(submittedPost)
}

func ReadPosts(res http.ResponseWriter, request *http.Request) {
	json.NewEncoder(res).Encode(posts)
}

func ReadPost(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "postId")
	postID, err := strconv.Atoi(id)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Invalid post ID!"))
		return
	}

	postWithID := posts[postID-1]
	json.NewEncoder(res).Encode(postWithID)
}

func UpdatePost(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "postId")
	postID, err := strconv.Atoi(id)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Invalid post Id !"))
		return
	}

	if postID > len(posts) || postID < 1 {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Post Id is out of range !"))
		return
	}

	postWithID := &posts[postID-1]
	err = json.NewDecoder(req.Body).Decode(postWithID)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Invalid data submitted !"))
		return
	}
	json.NewEncoder(res).Encode(postWithID)
}

func DeletePost(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "postId")
	postID, err := strconv.Atoi(id)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Something is wrong!"))
		return
	}

	if postID > len(posts) || postID < 1 {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Post Id is out of range !"))
		return
	}

	posts = append(posts[:postID-1], posts[postID:]...)
	res.Write([]byte("Post deleted successfully"))
}

//Create post, Read All Posts, Read Single Post, Update Post, Delete Post
