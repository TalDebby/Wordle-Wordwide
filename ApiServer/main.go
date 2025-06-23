package main

import (
	"fmt"
	"net/http"
	languagesCrud "server/crud"
	"server/middleware"
	"server/routes"
	languagesService "server/service"
)

func main() {
	fmt.Println("Starting Server")

	languagesRepository := languagesCrud.NewJsonWordsRepository()
	languagesService := languagesService.NewLanguagesService(languagesRepository)

	mux := http.NewServeMux()

	routes.AddRoutes(mux, languagesService)

	defaultMiddlewars := middleware.CreateStack(
		middleware.Logging,
		middleware.Recovery,
		middleware.CORS,
	)

	if err := http.ListenAndServe("localhost:8080", defaultMiddlewars(mux)); err != nil {
		fmt.Println(err.Error())
	}
}
