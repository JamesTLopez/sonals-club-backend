package controllers

import (
	"net/http"
	"sonalsguild/helpers"
	"sonalsguild/internal/services"
)

var samplesService services.Sample


func GetAllSamples(w http.ResponseWriter, r *http.Request) {
	_, err := samplesService.GetAllSamples()

	if(err != nil) {
		
		helpers.MessageLogs.ErrorLog.Println("ERROR BOI")
		helpers.ErrorJson(w, err)
		return
	}

	helpers.WriteJson(w, http.StatusOK, helpers.Envelop{"songs":"sample"})
}