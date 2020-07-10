package main

import (
	"github.com/caarlos0/env"
	"github.com/gorilla/mux"
	"github.com/igorkichuk/ciklum/internal/controller"
	"github.com/igorkichuk/ciklum/internal/provider"
	"github.com/igorkichuk/ciklum/internal/usecase"
	"github.com/joho/godotenv"
	"log"
	"net"
	"net/http"
)

type config struct {
	ArticleUrl string `env:"ARTICLE_URL,required"`
	CMUrl      string `env:"CONTENT_MARKETING_URL,required"`
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load config,e:", err.Error())
	}

	cfg := config{}
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatal("Failed to parse config,e:", err.Error())
	}

	ciklumProvider := provider.NewCiklumProvider()
	articleUsecase := usecase.NewArticleUsecase(ciklumProvider, cfg.ArticleUrl, cfg.CMUrl)
	api := controller.NewApiController(articleUsecase)

	r := mux.NewRouter()
	r.HandleFunc("/two_responses", api.TwoResponses)

	tcpAddr := net.TCPAddr{Port: 8080}
	log.Println("[INFO] Server is starting on port", 8080)
	if err := http.ListenAndServe(tcpAddr.String(), r); err != nil {
		log.Fatal(err)
	}
}
