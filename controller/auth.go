package controller

import (
	"encoding/json"
	"fmt"
	"goserver/middleware"
	"net/http"
)

type loginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func login(w http.ResponseWriter, r *http.Request) {
	body := r.Context().Value("body").(*loginRequest)
	fmt.Println(body)
	json.NewEncoder(w).Encode(body)
}

func status(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func HandleAuthRouter() *http.ServeMux {
	authRouter := http.NewServeMux()
	authRouter.HandleFunc("POST /login", middleware.BodyValidator[loginRequest](login))
	authRouter.HandleFunc("GET /status", status)
	return authRouter
}
