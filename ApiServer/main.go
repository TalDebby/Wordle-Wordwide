package main

import (
	"fmt"
	"net/http"
	languagesCrud "server/crud"
	"server/routes"
	languagesService "server/service"
)

func main() {
	fmt.Println("Starting Server")

	wordsRepository := languagesCrud.NewJsonWordsRepository()
	wordsService := languagesService.NewLanguagesService(wordsRepository)

	mux := http.NewServeMux()

	routes.AddRoutes(mux, wordsService)

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err.Error())
	}
}
