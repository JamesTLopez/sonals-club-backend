package controllers

import (
	"encoding/json"
	"net/http"
	"sonalsguild/helpers"
	"sonalsguild/internal/services"

	"github.com/go-chi/chi/v5"
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

// GET/songs/{id}
func GetSongById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r,"id")

	song, err := song.GetSongById(id)
	
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJson(w, http.StatusOK, song)
}


// POST/songs/song
func CreateSong(w http.ResponseWriter, r *http.Request) {
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

//PUT/songs/{id}
func UpdateSong(w http.ResponseWriter, r *http.Request) {
	var songData services.Song
	id := chi.URLParam(r,"id");

	err := json.NewDecoder(r.Body).Decode(&songData)

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	songUpdated, err := song.UpdateSong(id,songData)

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	helpers.WriteJson(w, http.StatusOK, songUpdated)
}