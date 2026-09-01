package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestRecodingWinsAndRetrievemThem(t *testing.T)  {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	assertScoreOfPlayers(t, response.Result().Status, "3")
}