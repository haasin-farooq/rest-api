package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Post struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}

var posts []Post

func main() {
	posts = append(posts, Post{Id: "1", Title: "Post 1", Body: "This is the first post."})

	router := mux.NewRouter()

	router.HandleFunc("/posts", createPost).Methods("POST")
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.Id = strconv.Itoa(rand.Intn(1000000))
	posts = append(posts, post)
	json.NewEncoder(w).Encode(&post)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range posts {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Post{})
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range posts {
		if item.Id == params["id"] {
			posts = append(posts[:i], posts[i+1:]...)
			
			var post Post
			_ = json.NewDecoder(r.Body).Decode(&post)
			post.Id = params["id"]
			posts = append(posts, post)
			json.NewEncoder(w).Encode(&post)

			return
		}
	}
	json.NewEncoder(w).Encode(posts)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range posts {
		if item.Id == params["id"] {
			posts = append(posts[:i], posts[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(posts)
}