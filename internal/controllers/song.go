package controllers

import (
	"encoding/json"
	"net/http"
	"sonalsguild/helpers"
	"sonalsguild/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
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
	userValue := r.Context().Value("user").(jwt.MapClaims)
	user_id, ok := userValue["spotify_id"].(string)

	if !ok {
		helpers.MessageLogs.ErrorLog.Println("Something went wrong when grabbing the id")

		return
	}
	
	
	err := json.NewDecoder(r.Body).Decode(&songData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	songCreated, err := song.CreateSong(user_id, songData)

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	// helpers.WriteJson(w, http.StatusOK, "success")
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


//DELETE/songs/{id}
func DeleteSong(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r,"id");

	err := song.DeleteSong(id)

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	helpers.WriteJson(w, http.StatusOK, "Song successfully deleted")
}