package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	languagesService "server/service"
)

func AddRoutes(mux *http.ServeMux, languagesService *languagesService.LanguagesService) {
	mux.HandleFunc("/languages/{language}/words/dailyword", func(w http.ResponseWriter, r *http.Request) {
		language := r.PathValue("language")
		fmt.Print("Language: ", language)
		if language == "" {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		word, err := languagesService.GetDailyWord(language)
		if err != nil {
			fmt.Println(err.Error())
			if err.Error() == "not found" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			return
		}

		dailyWordJson, _ := json.Marshal(map[string](string){"dailyword": word})

		fmt.Printf("%s\n", string(dailyWordJson))

		w.Header().Set("content-type", "application/json")
		w.Write(dailyWordJson)
	})

	mux.HandleFunc("GET /languages/{language}/words", func(w http.ResponseWriter, r *http.Request) {
		language := r.PathValue("language")
		if language == "" {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		words, err := languagesService.GetWords(language)
		if err != nil {
			fmt.Println(err.Error())
			if err.Error() == "not found" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			return
		}

		wordsJson, _ := json.Marshal(map[string]([]string){language: *words})

		w.Write(wordsJson)
	})

	mux.HandleFunc("GET /languages", func(w http.ResponseWriter, r *http.Request) {
		languages, err := languagesService.GetLanguages()
		if err != nil {
			fmt.Println(err.Error())
			if err.Error() == "not found" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			return
		}

		wordsJson, _ := json.Marshal(languages)

		w.Write(wordsJson)
	})
}
