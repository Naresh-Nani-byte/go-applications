package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayersScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	
	router := http.NewServeMux()
	

	router.Handle("/league", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	
	router.Handle("/players/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		player := strings.TrimPrefix(r.URL.Path, "/players/")

		switch r.Method {
		case http.MethodPost:
			p.processWins(w, player)
		case http.MethodGet:
			p.showPlayerScore(w, player)
	}
	}))
	router.ServeHTTP(w, r)
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) PlayerHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWins(w, player)
	case http.MethodGet:
		p.showPlayerScore(w, player)
	}
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux()
	router.Handle("/players/", http.HandlerFunc(p.PlayerHandler))
	router.Handle("/league/", http.HandlerFunc(p.leagueHandler))

	p.Handler = router

	return p
}

func (p *PlayerServer) showPlayerScore(w http.ResponseWriter, player string) {

	score := p.store.GetPlayersScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprint(w, score)
}


func (p *PlayerServer) processWins(w http.ResponseWriter, player string)  {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}





