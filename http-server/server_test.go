package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
	winCalls []string
}

func TestGETPlayer(t *testing.T)  {
	
	store := StubPlayerStore{
		map[string]int{
			"Naresh" : 20,
			"Shiva": 10,
		},
		nil,
	}

	server := &PlayerServer{&store}

	t.Run("Return the Naresh Score", func(t *testing.T) {
		request := newGetScoreRequest("Naresh")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)	
		
		assertStatus(t, response.Code, http.StatusOK)
		assertScoreOfPlayers(t, response.Body.String(), "20")
	})

	t.Run("Let's get the Shiva scor",  func(t *testing.T) {
		request := newGetScoreRequest("Shiva")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertScoreOfPlayers(t, response.Body.String(), "10")
	})

	t.Run("Returns 404 for missing players",  func(t *testing.T) {
		request := newGetScoreRequest("Marine")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
		
	})
}

func TestSoreWins(t *testing.T)  {
	store := StubPlayerStore{
		map[string]int{},
		nil,
	}
	
	server := &PlayerServer{&store}

	t.Run("Store the Naveen's score", func(t *testing.T) {
		player := "Naveen"
		request := newPostScoreRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusAccepted)
		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
		}
	})
}

func assertScoreOfPlayers(t testing.TB, got, want string)  {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
func newGetScoreRequest(name string) *http.Request  {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s",name), nil)
	return req
}

func newPostScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertStatus(t testing.TB, got, want int)  {
	t.Helper()
	if got != want {
		t.Errorf("response status is wrong, got %d want %d", got, want)
	}
}

func (s *StubPlayerStore) GetPlayersScore(name string) int {
		score := s.scores[name]
		return score
}

func (s *StubPlayerStore) RecordWin(name string)  {
	s.winCalls = append(s.winCalls, name)
}
