package handlers

import (
	"encoding/json"
	"net/http"
	"separation/auth/db"
	"separation/auth/dtos"
	"separation/auth/utils"
	"time"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var req dtos.SignUpRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req.Password, err = utils.HashPassword(req.Password)

	err = db.InsertIntoAuth(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func Validate(w http.ResponseWriter, r *http.Request) {
	var token string = r.Header.Get("Authorization")

	username, err := utils.IsValid(token, "access")
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(username)
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	refreshToken := r.Header.Get("Authorization")
	var accessToken string

	username, err := utils.IsValid(refreshToken, "refresh")
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	accessToken, err = utils.GenerateToken(username, "access", time.Minute*15)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req dtos.SignUpRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hash, err := db.SelectHash(req.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = utils.CheckPassword(hash, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	refreshToken, err := utils.GenerateToken(req.Username, "refresh", time.Hour*24*7)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	accessToken, err := utils.GenerateToken(req.Username, "access", time.Minute*15)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tokens := struct {
		RefreshToken string `json:"refreshToken"`
		AccessToken  string `json:"accessToken"`
	}{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tokens)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var token string = r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Нет заголовка с JWT", http.StatusUnauthorized)
		return
	}

	tokenUsername, err := utils.IsValid(token, "access")
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var deleteUsername string = r.URL.Path[len("/delete/"):]

	if tokenUsername != deleteUsername {
		http.Error(w, "Удалить можно только себя", http.StatusForbidden)
		return
	}

	err = db.DeleteFromAuth(deleteUsername)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
