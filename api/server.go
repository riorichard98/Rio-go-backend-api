package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Item struct{
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
}

type Server struct{
	*mux.Router
	shoppingItems []Item
}

func newServer() *Server{
	s := &Server{
		Router: mux.NewRouter(),
		shoppingItems: []Item{},
	}
	return s
}

func (s *Server) createShoppingItem () http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		items := []Item{}
		if err := json.NewDecoder(r.Body).Decode(&items) ; err != nil{
			http.Error(w,err.Error(),http.StatusBadRequest)
			return
		}
		s.shoppingItems = append(s.shoppingItems, Item{
			ID: uuid.New(),
		})
		w.Header().Set("Content-Type","application/json")
		if err := json.NewEncoder(w).Encode(items) ; err != nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) removeShoppingItem () http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		idStr,_ := mux.Vars(r)["id"]
		id,err := uuid.Parse(idStr)
		if err != nil{
			http.Error(w,err.Error(),http.StatusBadRequest)
		}

		for i ,item := range s.shoppingItems{
			if item.ID == id {
				s.shoppingItems = append(s.shoppingItems[:i],s.shoppingItems[i+1:]...)
				break
			}
		}
	}
}