package controllers

import (
	"condoapp/src/repositories"
	"condoapp/src/viewmodels"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func GetProperties(responseWriter http.ResponseWriter, request *http.Request) {
	properties := repositories.GetProperties()
	residents := repositories.GetResidents()

	for i := 0; i < len(properties); i++ {
		if properties[i].PropertyId == residents[i].PropertyId {
			properties[i].Residents = append(properties[i].Residents, residents[i])
		}
	}

	viewmodelData := viewmodels.PropertyViewViewModel {
		Properties: properties,
	}

	templateResponse, error := template.ParseFiles("./src/views/property-view.html")
	if error != nil {
		http.Error(responseWriter, error.Error(), http.StatusInternalServerError)
		return
	}
	error = templateResponse.Execute(responseWriter, viewmodelData)
	if error != nil {
		fmt.Println(error)
		responseWriter.Write([]byte(error.Error()))
	}
}

func Ping(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode("Pong")
}
