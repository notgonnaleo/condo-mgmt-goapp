package controllers

import (
	"condoapp/src/models"
	"encoding/json"
	"net/http"
)

func GetProperties(responseWriter http.ResponseWriter, request *http.Request) {
	properties := []models.Property{
		{PropertyId: 1, PropertyAddressName: "LEO ADDRESS"},
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(properties)
}
