package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/kpfaulkner/grafanadatadog/pkg/helpers"
	"github.com/kpfaulkner/grafanadatadog/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Server is the main struct for the datasource
type Server struct {
  dd *Datadog
  config models.Config
}

func NewServer() *Server {
	s := Server{}
	s.config = readConfig()
  s.dd = NewDatadog(s.config.DatadogAPIKey, s.config.DatadogAppKey)
	s.routes()

	return &s
}

func readConfig() models.Config{
	configFile, err := os.Open("config.json")
	defer configFile.Close()
	if err != nil {
		log.Panic("Unable to read config.json")
	}

	config := models.Config{}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config
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
	searchResults := []string{"test1", "test2"}

  j, err := json.Marshal(searchResults)
  if err != nil {
	  log.Printf("Error reading body: %v", err)
	  http.Error(w, "cant search", http.StatusBadRequest)
	  return
  }

	w.WriteHeader(http.StatusOK)
  w.Write(j)
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

  ddResponse, err := s.dd.queryDatadog("@environment:prod status:(error)", request.Range.From, request.Range.To)
  if err != nil {
	  log.Printf("unable to query datadog: %v", err)
	  http.Error(w, "Sorry, cannot query datadog", http.StatusInternalServerError)
	  return
  }

  response, err := helpers.ConvertDDResponseToGrafanaResponse( *ddResponse)
	if err != nil {
		log.Printf("unable to convert DD response: %v", err)
		http.Error(w, "Sorry, unable to process datadog response", http.StatusInternalServerError)
		return
	}

  fmt.Printf("query response %v\n", response)

  json, err := json.Marshal(response)
  if err != nil {
	  log.Printf("unable to generate response: %v", err)
	  http.Error(w, "Sorry, unable to generate datadog response", http.StatusInternalServerError)
	  return
  }

  w.Write( json )
	w.WriteHeader(http.StatusOK)
}

func (s *Server) annotations( w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *Server) Run() error {
	http.ListenAndServe(":8080", nil)
	return nil
}








