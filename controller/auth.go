package controller

import (
	"encoding/json"
	"fmt"
	"goserver/core"
	"goserver/middleware"
	"net/http"
)

type loginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type signupRequest struct {
	Username string `json:"username" validate:"required,gt=4"`
	Password string `json:"password" validate:"required,gt=6"`
	Email    string `json:"email" validate:"required,email"`
}

func login(w http.ResponseWriter, r *http.Request) {
	body := r.Context().Value("body").(*loginRequest)
	fmt.Println(body)
	jwt := core.NewJwtLib()
	token, refresh := jwt.SignJwt(core.StandardClaims{Id: "123"})
	jwt.ValidateJwt(token)
	json.NewEncoder(w).Encode(map[string]string{"token": token, "refresh": refresh})
}

func register(w http.ResponseWriter, r *http.Request) {
	body := r.Context().Value("body").(*signupRequest)
	fmt.Println(body)
}

func status(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func HandleAuthRouter() *http.ServeMux {
	authRouter := http.NewServeMux()
	authRouter.HandleFunc("POST /login", middleware.BodyValidator[loginRequest](login))
	authRouter.HandleFunc("POST /signup", middleware.BodyValidator[signupRequest](register))
	authRouter.HandleFunc("GET /status", status)
	return authRouter
}
