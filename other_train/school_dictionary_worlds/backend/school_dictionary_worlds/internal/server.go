package internal

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"log"
	"net/http"
	"school_dictionary_worlds/games"
)

func RunServer() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/cypher_coordinates", func(w http.ResponseWriter, r *http.Request) {
		resp := make(map[string]string)
		resp["message"] = "Status Created"

		exercise := games.LunchCryptographer()

		jsonResp, err := json.Marshal(exercise)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	fmt.Println("Listening on http://localhost:3000")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}

}
