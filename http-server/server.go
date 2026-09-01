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
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWins(w, player)
	case http.MethodGet:
		p.showPlayerScore(w, player)
	}
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





