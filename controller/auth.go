package controller

import (
	"encoding/json"
	"net/http"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login(w http.ResponseWriter, r *http.Request) {
	docoder := json.NewDecoder(r.Body)
	body := loginRequest{}
	if err := docoder.Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("ok"))
	defer r.Body.Close()
}

func status(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func HandleAuthRouter() *http.ServeMux {
	authRouter := http.NewServeMux()
	authRouter.HandleFunc("POST /login", login)
	authRouter.HandleFunc("GET /status", status)
	return authRouter
}
