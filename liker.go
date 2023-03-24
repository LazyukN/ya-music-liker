package main

import (
	"context"
	"fmt"
	"github.com/LazyukN/go-yamusic-extended/yamusic"
	"github.com/rubyist/circuitbreaker"
	"log"
	"net/http"
	"time"
)

// https://github.com/MarshalX/yandex-music-api/discussions/513
const token = "secret_token"
const playlistId = 100
const userId = 100

func main() {
	circuitClient := circuit.NewHTTPClient(time.Second*5, 10, nil)
	client := yamusic.NewClient(
		yamusic.HTTPClient(circuitClient),
		yamusic.AccessToken(userId, token),
	)

	playlist, resp, err := client.Playlists().Get(context.Background(), client.UserID(), playlistId)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal("http status is not 200")
	}
	var trackIds []string
	for i := 0; i < len(playlist.Result.Tracks); i++ {
		track := playlist.Result.Tracks[i]
		trackIds = append(trackIds, track.Track.ID)
	}

	like, resp, err := client.Likes().Like(context.Background(), "track", trackIds)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal("http status is not 200")
	}
	fmt.Println("Like: ", like.Result)

}
