package hive

import (
	"encoding/json"
	"io"
	"net/http"
)

// buildRequest builds a request with the given arguments.
func buildRequest(args ...string) string {
	var str = "https://api.playhive.com/v0/game"
	for _, arg := range args {
		str += "/" + arg
	}
	return str
}

// PlayerRequest is a struct used to create player requests.
type PlayerRequest struct {
	time   Time
	game   Game
	player string
}

// NewPlayerRequest creates a new *PlayerRequest.
func NewPlayerRequest(time Time, game Game, player string) *PlayerRequest {
	return &PlayerRequest{
		time:   time,
		game:   game,
		player: player,
	}
}

// Do does a request to hive and returns a PlayerResponse.
func (r *PlayerRequest) Do() (PlayerResponse, error) {
	var playerResponse PlayerResponse

	time := r.time.String()
	game := r.game.String()

	req := buildRequest(time, game, r.player)

	response, err := http.Get(req)
	if err != nil {
		return playerResponse, err
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return playerResponse, err
	}

	err = json.Unmarshal(data, &playerResponse)
	if err != nil {
		return PlayerResponse{}, err
	}
	playerResponse.FirstPlayedTime = firstPlayedTime(playerResponse)

	return playerResponse, nil
}
