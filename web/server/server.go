package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/edznux/lyfe/collectors"
	"github.com/edznux/lyfe/collectors/strava"
	"github.com/edznux/lyfe/storage"
	"github.com/edznux/lyfe/storage/sqliteStore"
	"github.com/spf13/viper"
)

type Server struct {
	Store     *storage.Store
	Collector collectors.Collector
}

func NewServer(store *storage.Store, collector collectors.Collector) *Server {
	return &Server{
		Store:     store,
		Collector: collector,
	}
}

func (s *Server) Start() {
	s.RegisterRoutes()

	addrPort := "127.0.0.1:8080"
	fmt.Println("Starting server on: http://" + addrPort)

	http.ListenAndServe(addrPort, nil)
}

func (s *Server) RegisterRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		scope := "activity:read_all,profile:read_all"
		clientId := viper.GetInt("strava.client_id")
		redirectURI := "http://127.0.0.1:8080/strava/callback"
		loginElement := fmt.Sprintf("<a href='https://www.strava.com/oauth/authorize?client_id=%d&response_type=code&scope=%s&redirect_uri=%s'>Strava login!</a>", clientId, scope, redirectURI)
		w.Write([]byte(loginElement))
	})

	http.HandleFunc("/collect/strava", func(w http.ResponseWriter, r *http.Request) {
		err := s.Collector.Collect(r.Context(), 0)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("ok"))
	})
	http.HandleFunc("/strava/callback", func(w http.ResponseWriter, r *http.Request) {

		code := r.URL.Query().Get("code")
		clientId := viper.GetInt("strava.client_id")
		clientSecret := viper.GetString("strava.client_secret")

		getTokenURL := fmt.Sprintf("https://www.strava.com/api/v3/oauth/token?client_id=%d&client_secret=%s&code=%s&grant_type=authorization_code", clientId, clientSecret, code)
		fmt.Println(getTokenURL)

		httpClient := http.Client{}
		req, err := http.NewRequest("POST", getTokenURL, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp, err := httpClient.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(string(body))

		var stravaResp strava.StravaOAuthResponse
		err = json.Unmarshal(body, &stravaResp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Printf("%+v", stravaResp)
		err = s.Store.Querier.CreateAuthToken(r.Context(), sqliteStore.CreateAuthTokenParams{
			AthleteID:    stravaResp.Athlete.ID,
			AccessToken:  stravaResp.AccessToken,
			RefreshToken: stravaResp.RefreshToken,
			ExpiresAt:    stravaResp.ExpiresAt,
			Firstname:    stravaResp.Athlete.Firstname,
			Lastname:     sql.NullString{String: stravaResp.Athlete.Lastname, Valid: true},
			ExpiresIn:    stravaResp.ExpiresIn,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("ok"))
	})
}
