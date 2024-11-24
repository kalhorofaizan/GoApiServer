package main

import (
	"github.com/joho/godotenv"
	"goserver/controller"
	"goserver/middleware"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file " + err.Error())
		return
	}
	middlewareChain := middleware.HandleChainMiddleware(
		middleware.EnableCors,
		middleware.LogApi)
	mainRouter := http.NewServeMux()
	apiRouter := http.NewServeMux()
	apiRouter.Handle("/auth/", http.StripPrefix("/auth", controller.HandleAuthRouter()))
	mainRouter.Handle("/api/", http.StripPrefix("/api", apiRouter))
	println(":" + os.Getenv("PORT"))
	server := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: middlewareChain(mainRouter),
	}
	log.Fatal(server.ListenAndServe())
}
