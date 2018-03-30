package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"../models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

var authorsDAO = models.AuthorsDAO{Collection: "authors"}
var storiesDAO = models.StoriesDAO{Collection: "stories"}
var commentsDAO = models.CommentsDAO{Collection: "comments"}
var db *mgo.Database

//Router to serve API requests
var Router = mux.NewRouter()

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	connect()
	Router.HandleFunc("/authors", func(w http.ResponseWriter, r *http.Request) {
		configResponse(&w)
		findAll(authorsDAO, w, r)
	}).Methods("GET")

	Router.HandleFunc("/stories", func(w http.ResponseWriter, r *http.Request) {
		configResponse(&w)
		findAll(storiesDAO, w, r)
	}).Methods("GET")

	Router.HandleFunc("/comments", func(w http.ResponseWriter, r *http.Request) {
		configResponse(&w)
		findAll(commentsDAO, w, r)
	}).Methods("GET")
}

func findAll(dao models.ModelDAO, w http.ResponseWriter, r *http.Request) {
	var collection interface{}
	switch dao.(type) {
	case models.AuthorsDAO:
		if authors, err := authorsDAO.FindAll(db); handleErr(err, "Could not retrieve authors") {
			collection = authors
		}
	case models.StoriesDAO:
		if stories, err := storiesDAO.FindAll(db); handleErr(err, "Could not retrieve stories") {
			collection = stories
		}
	case models.CommentsDAO:
		if comments, err := commentsDAO.FindAll(db); handleErr(err, "Could not retrieve comments") {
			collection = comments
		}
	default:
		log.Fatalf("%v is an unrecognized DAO", dao)
	}
	json.NewEncoder(w).Encode(collection)
}
