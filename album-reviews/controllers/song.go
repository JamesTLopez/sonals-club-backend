package controllers

import (
	"encoding/json"
	"net/http"
	"sonalsguild/helpers"
	"sonalsguild/services"
)


var song services.Song

// GET/songs

func GetAllSongs(w http.ResponseWriter, r *http.Request) {
	all, err := song.GetAllSongs()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	helpers.WriteJson(w, http.StatusOK, helpers.Envelop{"songs":all})
}

// POST/songs/song
func CreateSongs(w http.ResponseWriter, r *http.Request) {
	var songData services.Song
	err := json.NewDecoder(r.Body).Decode(&songData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	helpers.WriteJson(w, http.StatusOK, songData)

	songCreated, err := song.CreateSong(songData)

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJson(w, http.StatusOK, songCreated)
}