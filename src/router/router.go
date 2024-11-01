package router

import (
	"condoapp/src/controllers"
	"net/http"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/properties", controllers.GetProperties)
	mux.HandleFunc("/ping", controllers.Ping)
	return mux
}
