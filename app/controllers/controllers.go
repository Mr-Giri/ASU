package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	types "music/web-music-microservice/app/types"

	"github.com/gin-gonic/gin"
)

const (
	BASE_URL = "https://api.musixmatch.com/ws/1.1/"

	API_KEY = "aa55d46fac7bd1674d6909a7f6446e29" //we should do os.getenv but i have hardcoded it
	PAGE    = 1
)

func GetTracksByRegion(c *gin.Context) {
	regionName := c.Param("name")

	response, err := http.Get(BASE_URL + "chart.tracks.get?apikey=" + API_KEY + "&page=1&page_size=5&country=" + regionName + "&f_has_lyrics=1")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	responseJSON := types.TrackRegionResponse{}
	err = json.Unmarshal(body, &responseJSON)
	if err != nil {
		panic(err.Error())
	}

	ResponseOf := types.TrackRegionResponse{}

	for i, v := range responseJSON.Message.Body.TrackList {
		ResponseOf.Message.Header = responseJSON.Message.Header
		lyrics, err := GetLyricsByTrackId(fmt.Sprintf("%v", v.Track.TrackId))
		if err != nil {
			fmt.Println(err)
		}

		artist, err := GetArtistById(fmt.Sprintf("%v", v.Track.ArtistId))
		if err != nil {
			fmt.Println(err)
		}

		ResponseOf.Message.Body = responseJSON.Message.Body
		ResponseOf.Message.Body.TrackList[i].Track.Lyrics = lyrics.Message.Body.Lyrics
		ResponseOf.Message.Body.TrackList[i].Track.Artist = artist.Message.Body.Artist
	}

	c.JSON(http.StatusOK, ResponseOf)
}

func GetLyricsByTrackId(trackId string) (types.LyricsResponse, error) {
	response, err := http.Get(BASE_URL + "track.lyrics.get?track_id=" + trackId + "&apikey=" + API_KEY)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	lyricsResp := types.LyricsResponse{}
	err = json.Unmarshal(body, &lyricsResp)
	if err != nil {
		return types.LyricsResponse{}, err
	}

	return lyricsResp, nil
}

func GetArtistById(artistId string) (types.ArtistResponse, error) {
	response, err := http.Get(BASE_URL + "artist.get?artist_id=" + artistId + "&apikey=" + API_KEY)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	artistResp := types.ArtistResponse{}
	err = json.Unmarshal(body, &artistResp)
	if err != nil {
		return types.ArtistResponse{}, err
	}

	return artistResp, nil
}
