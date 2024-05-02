package controllers

import (
	"encoding/json"
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

func CreateSample(w http.ResponseWriter, r *http.Request){
	var sampleData services.Sample
	err := json.NewDecoder(r.Body).Decode(&sampleData)

	if err != nil {
		helpers.ErrorJson(w,err)
		return 
	}

	sampleCreated, err := samplesService.CreateSample(sampleData)

	if err != nil {
		helpers.ErrorJson(w,err)
		return
	}

	helpers.WriteJson(w, http.StatusOK, sampleCreated)

}