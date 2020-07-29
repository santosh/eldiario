package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/santosh/eldiario/handler"
)

var err error

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {
	// connection string has to match to the container name in docker-compose
	session, err := mgo.Dial("eldiario_mongo")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	ensureIndex(session)

	r := mux.NewRouter()
	r.HandleFunc("/entry", handler.GetEntries(session))
	r.HandleFunc("/entry", handler.CreateEntry(session))
	r.HandleFunc("/entry/:id", handler.GetEntry(session))
	r.HandleFunc("/entry/:id", handler.UpdateEntry(session))
	r.HandleFunc("/entry/:id", handler.DeleteEntry(session))

	fmt.Println("Listening at http://127.0.0.1:8080")
	http.ListenAndServe(":8080", r)
}

func ensureIndex(s *mgo.Session) {
	session := s.Copy()
	defer session.Close()

	c := session.DB("eldiario").C("entries")

	index := mgo.Index{
		// TODO: Can we please use mongo's native _id?
		Key:        []string{"id"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}

	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}
