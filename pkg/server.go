package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/kpfaulkner/grafanadatadog/pkg/helpers"
	"github.com/kpfaulkner/grafanadatadog/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
)

// Server is the main struct for the datasource
type Server struct {
  dd *Datadog
}

func NewServer() *Server {
	s := Server{}
  s.dd = NewDatadog()
	s.routes()
	return &s
}

// routes sets up routes... duh :)
func (s *Server) routes() {
	http.HandleFunc("/", s.testConnection)
	http.HandleFunc("/search", s.search)
	http.HandleFunc("/query", s.query)
	http.HandleFunc("/annotations", s.annotations)
}

func (s *Server) testConnection( w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// search is a POST
func (s *Server) search( w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

}

// query is a POST
func (s *Server) query( w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	var request models.Request
  err = json.Unmarshal(body, &request)
  if err != nil {
	  log.Printf("Error unmarshalling body: %v", err)
	  http.Error(w, "can't read body", http.StatusBadRequest)
	  return
  }

  ddResponse, err := s.dd.queryDatadog("fluffy")
  if err != nil {
	  log.Printf("unable to query datadog: %v", err)
	  http.Error(w, "Sorry, cannot query datadog", http.StatusInternalServerError)
  }

  response := helpers.ConvertDDResponseToGrafanaResponse( *ddResponse)
  fmt.Printf("query response %v\n", response)
	w.WriteHeader(http.StatusOK)
}

func (s *Server) annotations( w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *Server) Run() error {
	http.ListenAndServe(":8080", nil)
	return nil
}








