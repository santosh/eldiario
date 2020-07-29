package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/santosh/eldiario/model"
	"goji.io/pat"
)

// ErrorWithJSON represents errors in JSON for HTTP response
func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}\n", message)
}

// ResponseWithJSON returns HTTP response in JSON
func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
	w.Write([]byte("\n"))
}

// GetEntries fetches all the Entry stored
// TODO: Needs to throttle the count.
func GetEntries(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		c := session.DB("eldiario").C("entries")

		var entries []model.Entry
		err := c.Find(bson.M{}).All(&entries)
		if err != nil {
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)

			log.Println("Failed to get all entries:", err)
			return
		}

		respBody, err := json.MarshalIndent(entries, "", " ")
		if err != nil {
			log.Fatal(err)
		}

		ResponseWithJSON(w, respBody, http.StatusOK)
	}
}

// CreateEntry makes a new Entry
func CreateEntry(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		var entry model.Entry
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&entry)
		if err != nil {
			ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
			return
		}

		c := session.DB("eldiario").C("entries")

		err = c.Insert(entry)
		if err != nil {
			if mgo.IsDup(err) {
				ErrorWithJSON(w, "Entry with this ID already exists", http.StatusBadRequest)
				return
			}

			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed insert entry:", err)
			return
		}

		w.Header().Set("Content-Type", "application")
		w.Header().Set("Location", r.URL.Path+"/"+entry.ID)
		w.WriteHeader(http.StatusCreated)
	}
}

// GetEntry fetches a single Entry given it's ID
func GetEntry(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		id := pat.Param(r, "id")

		c := session.DB("eldiario").C("entries")

		var entry model.Entry
		err := c.Find(bson.M{"id": id}).One(&entry)
		if err != nil {
			// TODO: Return some meaningful response with some meaningful status code
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed to find entry: ", err)
			return
		}

		if entry.ID == "" {
			ErrorWithJSON(w, "Entry not found", http.StatusNotFound)
			return
		}

		respBody, err := json.MarshalIndent(entry, "", " ")
		if err != nil {
			log.Fatal(err)
		}

		ResponseWithJSON(w, respBody, http.StatusOK)
	}
}

// UpdateEntry updates data in mongodb for specific ID
func UpdateEntry(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		id := pat.Param(r, "id")

		var entry model.Entry
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&entry)
		if err != nil {
			ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
			return
		}

		c := session.DB("eldiario").C("entries")

		err = c.Update(bson.M{"id": id}, &entry)
		if err != nil {
			switch err {
			default:
				ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
				log.Println("Failed update entry: ", err)
				return
			case mgo.ErrNotFound:
				ErrorWithJSON(w, "Entry not found", http.StatusNotFound)
				return
			}
		}

		w.WriteHeader(http.StatusNoContent)

	}
}

// DeleteEntry removes a entry a row from db
func DeleteEntry(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		id := pat.Param(r, "id")

		c := session.DB("eldiario").C("entries")

		err := c.Remove(bson.M{"id": id})

		if err != nil {
			switch err {
			default:
				ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
				log.Println("Failed to delete entry: ", err)
				return
			case mgo.ErrNotFound:
				ErrorWithJSON(w, "Entry not found", http.StatusNotFound)
				return
			}
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
